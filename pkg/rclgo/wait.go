package rclgo

/*
#include <rcl/wait.h>
#include <rcl_action/wait.h>
*/
import "C"

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// DefaultWaitTimeout bounds each underlying rcl_wait call made by
// WaitSet.Run when the WaitSet's WaitTimeout is zero.
//
// A finite timeout is a safety net against lost wakeups in the RMW layer:
// rmw_zenoh_cpp 0.6.7 can erase a wakeup that races with wait-set attachment
// (ros2/rmw_zenoh#1032), leaving a request or reply queued while rmw_wait
// sleeps. With no timers in the wait set, rcl_wait passes an infinite timeout
// to the RMW and that sleep never ends. Bounding it caps the damage at one
// timeout period.
var DefaultWaitTimeout = 100 * time.Millisecond

// LostWakeupThreshold is how quickly rcl_wait must return ready, immediately
// after a timed-out wait, for Run to count the iteration as a recovered lost
// wakeup. When rmw_wait finds data already queued it returns without sleeping
// (microseconds); a message that genuinely arrives after the timeout lands at a
// uniformly random point in the next wait, so with the default values a false
// positive needs it to arrive in the first 0.5 % of the period.
var LostWakeupThreshold = 500 * time.Microsecond

type singleUse atomic.Bool

func (s *singleUse) reserve() bool {
	return (*atomic.Bool)(s).CompareAndSwap(false, true)
}

func (s *singleUse) release() {
	(*atomic.Bool)(s).Store(false)
}

// ErrWaitSetClosed is returned by WaitSet.Call and WaitSet.Run after the wait
// set has been closed.
var ErrWaitSetClosed = errors.New("wait set is closed")

// WaitSet waits on ROS entities and dispatches their callbacks from Run.
//
// The Add methods and Call are safe to use from any goroutine, including while
// Run is active and from callbacks dispatched by Run.
type WaitSet struct {
	rosID
	// WaitTimeout bounds each underlying rcl_wait call made by Run. Zero
	// means DefaultWaitTimeout; a negative value waits forever, which is the
	// historical behavior and leaves Run exposed to lost RMW wakeups.
	WaitTimeout time.Duration

	// The entity slices are only mutated under mu, and only by the Run
	// goroutine while running is set, so Run reads them without the lock.
	subscriptions   []*Subscription
	timers          []*Timer
	services        []*Service
	clients         []*Client
	actionClients   []*ActionClient
	actionServers   []*ActionServer
	guardConditions []*guardCondition

	mu      sync.Mutex
	idle    *sync.Cond // broadcast when inline drops to zero
	running bool
	closed  bool
	inline  int      // Call functions executing outside Run
	queue   []func() // work for Run to apply at the top of its next iteration

	rcl_wait_set_t C.rcl_wait_set_t
	cancelWait     *guardCondition
	wake           *guardCondition // wakes rcl_wait to apply queued work
	context        *Context
	lostWakeups    atomic.Uint64
}

func NewWaitSet() (*WaitSet, error) {
	if defaultContext == nil {
		return nil, errInitNotCalled
	}
	return defaultContext.NewWaitSet()
}

func (c *Context) NewWaitSet() (ws *WaitSet, err error) {
	const (
		subscriptionsCount   = 0
		guardConditionsCount = 0
		timersCount          = 0
		clientsCount         = 0
		servicesCount        = 0
		eventsCount          = 0
	)
	ws = &WaitSet{
		context:        c,
		rcl_wait_set_t: C.rcl_get_zero_initialized_wait_set(),
	}
	ws.idle = sync.NewCond(&ws.mu)
	defer onErr(&err, ws.Close)
	var rc C.rcl_ret_t = C.rcl_wait_set_init(
		&ws.rcl_wait_set_t,
		subscriptionsCount,
		guardConditionsCount,
		timersCount,
		clientsCount,
		servicesCount,
		eventsCount,
		c.rcl_context_t,
		*c.rcl_allocator_t,
	)
	if rc != C.RCL_RET_OK {
		return nil, errorsCast(rc)
	}
	ws.cancelWait, err = c.newGuardCondition()
	if err != nil {
		return nil, err
	}
	ws.wake, err = c.newGuardCondition()
	if err != nil {
		return nil, err
	}
	ws.addGuardConditions(ws.cancelWait, ws.wake)
	c.addResource(ws)
	return ws, nil
}

// Context returns the context s belongs to.
func (w *WaitSet) Context() *Context {
	return w.context
}

// AddSubscriptions adds subs to the wait set. While Run is active the
// subscriptions are picked up at the start of its next iteration.
func (w *WaitSet) AddSubscriptions(subs ...*Subscription) {
	subs = slices.Clone(subs)
	w.add(func() { w.subscriptions = append(w.subscriptions, subs...) })
}

// AddTimers adds timers to the wait set. See AddSubscriptions.
func (w *WaitSet) AddTimers(timers ...*Timer) {
	timers = slices.Clone(timers)
	w.add(func() { w.timers = append(w.timers, timers...) })
}

// AddServices adds services to the wait set. See AddSubscriptions.
func (w *WaitSet) AddServices(services ...*Service) {
	services = slices.Clone(services)
	w.add(func() { w.services = append(w.services, services...) })
}

// AddClients adds clients to the wait set. See AddSubscriptions.
func (w *WaitSet) AddClients(clients ...*Client) {
	clients = slices.Clone(clients)
	w.add(func() { w.clients = append(w.clients, clients...) })
}

// AddActionServers adds action servers to the wait set. See AddSubscriptions.
func (w *WaitSet) AddActionServers(servers ...*ActionServer) {
	servers = slices.Clone(servers)
	w.add(func() { w.actionServers = append(w.actionServers, servers...) })
}

// AddActionClients adds action clients to the wait set. See AddSubscriptions.
func (w *WaitSet) AddActionClients(clients ...*ActionClient) {
	clients = slices.Clone(clients)
	w.add(func() { w.actionClients = append(w.actionClients, clients...) })
}

func (w *WaitSet) addGuardConditions(guardConditions ...*guardCondition) {
	w.guardConditions = append(w.guardConditions, guardConditions...)
}

// add applies an entity-slice mutation now if Run is not active, otherwise
// queues it for Run. It never blocks on Run, so callbacks may use it.
func (w *WaitSet) add(apply func()) {
	w.mu.Lock()
	if !w.running {
		apply()
		w.mu.Unlock()
		return
	}
	w.queue = append(w.queue, func() {
		w.mu.Lock()
		apply()
		w.mu.Unlock()
	})
	w.mu.Unlock()
	// A failed trigger only delays the add to the next WaitTimeout expiry.
	_ = w.wake.Trigger()
}

/*
Call runs fn on the goroutine that owns the wait set's entities and returns
fn's error.

While Run is active, fn is queued, rcl_wait is woken, and Call blocks until Run
has executed fn at the start of its next iteration (or while exiting). When Run
is not active, fn runs inline on the calling goroutine and Run does not start
until it returns. After Close, Call returns ErrWaitSetClosed without running
fn.

Calling Call synchronously from a callback dispatched by the same wait set
deadlocks, because Run cannot execute fn until the callback returns. The Add
methods do not block and are safe to use from callbacks.
*/
func (w *WaitSet) Call(fn func() error) error {
	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		return ErrWaitSetClosed
	}
	if !w.running {
		w.inline++
		w.mu.Unlock()
		defer w.endInline()
		return fn()
	}
	done := make(chan error, 1)
	w.queue = append(w.queue, func() { done <- fn() })
	w.mu.Unlock()
	// A failed trigger only delays fn to the next WaitTimeout expiry or Run exit.
	_ = w.wake.Trigger()
	return <-done
}

func (w *WaitSet) endInline() {
	w.mu.Lock()
	w.inline--
	if w.inline == 0 {
		w.idle.Broadcast()
	}
	w.mu.Unlock()
}

// startRunning marks Run active once no inline Call is executing.
func (w *WaitSet) startRunning() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	for w.inline > 0 {
		w.idle.Wait()
	}
	switch {
	case w.closed:
		return ErrWaitSetClosed
	case w.running:
		return errors.New("wait set is already running")
	}
	w.running = true
	return nil
}

// stopRunning marks Run inactive and executes work queued too late for Run,
// so no Call blocks forever.
func (w *WaitSet) stopRunning() {
	w.mu.Lock()
	w.running = false
	pending := w.queue
	w.queue = nil
	w.inline++
	w.mu.Unlock()
	defer w.endInline()
	for _, fn := range pending {
		fn()
	}
}

// applyQueued executes queued work until the queue is empty, since queued
// functions may queue more. It reports whether anything ran.
func (w *WaitSet) applyQueued() bool {
	applied := false
	for {
		w.mu.Lock()
		pending := w.queue
		w.queue = nil
		w.mu.Unlock()
		if len(pending) == 0 {
			return applied
		}
		applied = true
		for _, fn := range pending {
			fn()
		}
	}
}

// reservedEntities tracks which entities a Run has reserved so entities added
// while running are reserved once and all are released on exit.
type reservedEntities struct {
	subscriptions, timers, services, clients int
	actionClients, actionServers, guardConds int
	reserved                                 []*singleUse
}

func reserveTail[T any](r *reservedEntities, entities []T, seen *int, waitable func(T) *singleUse) {
	for _, e := range entities[*seen:] {
		if s := waitable(e); s.reserve() {
			r.reserved = append(r.reserved, s)
		}
	}
	*seen = len(entities)
}

func (w *WaitSet) reserveNew(r *reservedEntities) {
	reserveTail(r, w.subscriptions, &r.subscriptions, func(s *Subscription) *singleUse { return &s.waitable })
	reserveTail(r, w.timers, &r.timers, func(t *Timer) *singleUse { return &t.waitable })
	reserveTail(r, w.services, &r.services, func(s *Service) *singleUse { return &s.waitable })
	reserveTail(r, w.clients, &r.clients, func(c *Client) *singleUse { return &c.waitable })
	reserveTail(r, w.actionClients, &r.actionClients, func(c *ActionClient) *singleUse { return &c.waitable })
	reserveTail(r, w.actionServers, &r.actionServers, func(s *ActionServer) *singleUse { return &s.waitable })
	reserveTail(r, w.guardConditions, &r.guardConds, func(g *guardCondition) *singleUse { return &g.waitable })
}

func (r *reservedEntities) releaseAll() {
	for _, s := range r.reserved {
		s.release()
	}
}

func (w *WaitSet) addResources(res *rosResourceStore) {
	for _, res := range res.resources {
		switch res := res.(type) {
		case *Subscription:
			w.AddSubscriptions(res)
		case *Timer:
			w.AddTimers(res)
		case *Service:
			w.AddServices(res)
		case *Client:
			w.AddClients(res)
		case *ActionServer:
			w.AddActionServers(res)
		case *ActionClient:
			w.AddActionClients(res)
		case *guardCondition: // Guard conditions are handled specially
		case *Node:
			w.addResources(&res.rosResourceStore)
		}
	}
}

// LostWakeupsRecovered returns how many times Run found work waiting
// immediately after a timed-out wait, the signature of a wakeup the RMW layer
// lost and the finite WaitTimeout recovered. Each occurrence is also logged
// at WARN level with the phrase "recovered lost wakeup".
func (w *WaitSet) LostWakeupsRecovered() uint64 {
	return w.lostWakeups.Load()
}

// waitTimeoutNanos resolves WaitTimeout to the rcl_wait timeout argument:
// -1 waits forever, anything else is a duration in nanoseconds.
func (w *WaitSet) waitTimeoutNanos() int64 {
	switch {
	case w.WaitTimeout < 0:
		return -1
	case w.WaitTimeout == 0:
		return DefaultWaitTimeout.Nanoseconds()
	default:
		return w.WaitTimeout.Nanoseconds()
	}
}

// recoveredLostWakeup reports whether a ready return from rcl_wait that took
// elapsed, directly after an iteration that timed out, should be counted as a
// recovered lost wakeup. See LostWakeupThreshold.
func recoveredLostWakeup(prevTimedOut bool, elapsed time.Duration) bool {
	return prevTimedOut && elapsed < LostWakeupThreshold
}

func (w *WaitSet) entityCounts() string {
	return fmt.Sprintf(
		"subscriptions=%d timers=%d services=%d clients=%d action_servers=%d action_clients=%d",
		len(w.subscriptions), len(w.timers), len(w.services), len(w.clients),
		len(w.actionServers), len(w.actionClients),
	)
}

/*
Run causes the current goroutine to block on this given WaitSet.
WaitSet executes the given timers and subscriptions and calls their callbacks on new events.

Each underlying rcl_wait is bounded by WaitTimeout (DefaultWaitTimeout when
zero) so that a wakeup lost inside the RMW layer delays the affected entity by
at most one period instead of hanging Run. A ready return that follows a
timed-out wait within LostWakeupThreshold is logged as a recovered lost wakeup
and counted in LostWakeupsRecovered.
*/
func (w *WaitSet) Run(ctx context.Context) (err error) {
	if ctx == nil {
		return errors.New("context must not be nil")
	}
	if err := w.startRunning(); err != nil {
		return err
	}
	var reserved reservedEntities
	defer func() {
		reserved.releaseAll()
		w.stopRunning()
	}()
	w.reserveNew(&reserved)
	errs := make(chan error, 1)
	defer func() {
		err = errors.Join(err, <-errs)
	}()
	errctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		defer close(errs)
		<-errctx.Done()
		errs <- w.cancelWait.Trigger()
	}()
	timeoutNanos := w.waitTimeoutNanos()
	var logger *Logger
	prevTimedOut := false
	for {
		// Queued work is applied only here: dispatch below indexes the rcl
		// arrays by the entity counts initEntities registered.
		if w.applyQueued() {
			w.reserveNew(&reserved)
		}
		if err := w.initEntities(); err != nil {
			return err
		}
		started := time.Now()
		rc := C.rcl_wait(&w.rcl_wait_set_t, C.int64_t(timeoutNanos))
		elapsed := time.Since(started)

		switch rc {
		case C.RCL_RET_OK:
			// fall through to handle the ready items
		case C.RCL_RET_TIMEOUT:
			prevTimedOut = true
			continue
		default:
			return errorsCast(rc)
		}

		guardConditions := unsafe.Slice(w.rcl_wait_set_t.guard_conditions, len(w.guardConditions))
		woken := false
		for i := range w.guardConditions {
			switch guardConditions[i] {
			case w.cancelWait.rclGuardCondition:
				return ctx.Err()
			case w.wake.rclGuardCondition:
				woken = true
			}
		}

		// A wake trigger is an expected immediate return, not a lost wakeup.
		if !woken && recoveredLostWakeup(prevTimedOut, elapsed) {
			w.lostWakeups.Add(1)
			if logger == nil {
				logger = GetLogger("rclgo.wait_set")
			}
			logger.Warn(fmt.Sprintf(
				"recovered lost wakeup: rcl_wait found work %v after a %v timeout (%s)",
				elapsed, time.Duration(timeoutNanos), w.entityCounts(),
			))
		}
		prevTimedOut = false
		timers := unsafe.Slice(w.rcl_wait_set_t.timers, len(w.timers))
		for i, t := range w.timers {
			if timers[i] != nil {
				t.Reset() //nolint:errcheck
				t.Callback(t)
			}
		}
		subs := unsafe.Slice(w.rcl_wait_set_t.subscriptions, len(w.subscriptions))
		for i, s := range w.subscriptions {
			if subs[i] != nil {
				s.Callback(s)
			}
		}
		svcs := unsafe.Slice(w.rcl_wait_set_t.services, len(w.services))
		for i, s := range w.services {
			if svcs[i] != nil {
				s.handleRequest()
			}
		}
		clients := unsafe.Slice(w.rcl_wait_set_t.clients, len(w.clients))
		for i, c := range w.clients {
			if clients[i] != nil {
				c.sender.HandleResponse()
			}
		}
		for _, s := range w.actionServers {
			s.handleReadyEntities(ctx, w)
		}
		for _, c := range w.actionClients {
			c.handleReadyEntities(w)
		}
	}
}

func (w *WaitSet) initEntities() error {
	if !C.rcl_wait_set_is_valid(&w.rcl_wait_set_t) {
		return errorsCastC(C.RCL_RET_WAIT_SET_INVALID, fmt.Sprintf("rcl_wait_set_is_valid() failed for wait_set='%v'", w))
	}
	var rc C.rcl_ret_t = C.rcl_wait_set_clear(&w.rcl_wait_set_t)
	if rc != C.RCL_RET_OK {
		return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_clear() failed for wait_set='%v'", w))
	}
	rc = C.rcl_wait_set_resize(
		&w.rcl_wait_set_t,
		C.size_t(len(w.subscriptions)+2*len(w.actionClients)),
		C.size_t(len(w.guardConditions)),
		C.size_t(len(w.timers)+len(w.actionServers)),
		C.size_t(len(w.clients)+3*len(w.actionClients)),
		C.size_t(len(w.services)+3*len(w.actionServers)),
		w.rcl_wait_set_t.size_of_events,
	)
	if rc != C.RCL_RET_OK {
		return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_resize() failed for wait_set='%v'", w))
	}
	for _, sub := range w.subscriptions {
		rc = C.rcl_wait_set_add_subscription(&w.rcl_wait_set_t, sub.rcl_subscription_t, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_subscription() failed for wait_set='%v'", w))
		}
	}
	for _, timer := range w.timers {
		rc = C.rcl_wait_set_add_timer(&w.rcl_wait_set_t, timer.rcl_timer_t, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_timer() failed for wait_set='%v'", w))
		}
	}
	for _, service := range w.services {
		rc = C.rcl_wait_set_add_service(&w.rcl_wait_set_t, service.rclService, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_service() failed for wait_set='%v'", w))
		}
	}
	for _, client := range w.clients {
		rc = C.rcl_wait_set_add_client(&w.rcl_wait_set_t, client.rclClient, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_client() failed for wait_set='%v'", w))
		}
	}
	for _, guardCondition := range w.guardConditions {
		rc = C.rcl_wait_set_add_guard_condition(&w.rcl_wait_set_t, guardCondition.rclGuardCondition, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_guard_condition() failed for wait_set='%v'", w))
		}
	}
	for _, server := range w.actionServers {
		rc = C.rcl_action_wait_set_add_action_server(&w.rcl_wait_set_t, &server.rclServer, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_action_server() failed for wait_set='%v'", w))
		}
	}
	for _, client := range w.actionClients {
		// rcl_action_wait_set_add_action_client is not thread-safe and reads the
		// client handle, which other goroutines may be querying through
		// ActionClient.IsServerAvailable; hold the client's handle lock for the
		// registration the same way the take path does.
		client.rclClientMu.Lock()
		rc = C.rcl_action_wait_set_add_action_client(&w.rcl_wait_set_t, &client.rclClient, nil, nil)
		client.rclClientMu.Unlock()
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_action_client() failed for wait_set='%v'", w))
		}
	}
	return nil
}

/*
Close frees the allocated memory
*/
func (w *WaitSet) Close() (err error) {
	if w.context == nil {
		return closeErr("wait set")
	}
	w.mu.Lock()
	w.closed = true
	w.mu.Unlock()
	w.context.removeResource(w)
	w.context = nil
	rc := C.rcl_wait_set_fini(&w.rcl_wait_set_t)
	if rc != C.RCL_RET_OK {
		err = errors.Join(err, errorsCast(rc))
	}
	var closeError closeError
	for _, gc := range []*guardCondition{w.cancelWait, w.wake} {
		if gc == nil {
			continue
		}
		gcErr := gc.Close()
		if gcErr != nil && !errors.As(gcErr, &closeError) {
			err = errors.Join(err, gcErr)
		}
	}
	return err
}
