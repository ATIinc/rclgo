package rclgo

/*
#include <stdlib.h>
#include <rcl/wait.h>
#include <rcl_action/wait.h>

// Handles of the entities that rclgo_wait_set_fill adds, in this order:
// subscriptions, timers, services, clients, guard conditions.
typedef struct {
	void ** handles;
	size_t subscriptions, timers, services, clients, guard_conditions;
} rclgo_wait_handles_t;

enum {
	RCLGO_FILL_VALID,
	RCLGO_FILL_CLEAR,
	RCLGO_FILL_RESIZE,
	RCLGO_FILL_SUBSCRIPTION,
	RCLGO_FILL_TIMER,
	RCLGO_FILL_SERVICE,
	RCLGO_FILL_CLIENT,
	RCLGO_FILL_GUARD_CONDITION,
};

// Clears, resizes, and refills ws from h in one cgo call. The extra_* sizes
// reserve room for action entities, which the caller adds afterwards.
static rcl_ret_t rclgo_wait_set_fill(
	rcl_wait_set_t * ws, const rclgo_wait_handles_t * h,
	size_t extra_subscriptions, size_t extra_timers, size_t extra_clients,
	size_t extra_services, int * failed)
{
	if (!rcl_wait_set_is_valid(ws)) {
		*failed = RCLGO_FILL_VALID;
		return RCL_RET_WAIT_SET_INVALID;
	}
	rcl_ret_t rc = rcl_wait_set_clear(ws);
	if (rc != RCL_RET_OK) {
		*failed = RCLGO_FILL_CLEAR;
		return rc;
	}
	rc = rcl_wait_set_resize(
		ws,
		h->subscriptions + extra_subscriptions,
		h->guard_conditions,
		h->timers + extra_timers,
		h->clients + extra_clients,
		h->services + extra_services,
		ws->size_of_events);
	if (rc != RCL_RET_OK) {
		*failed = RCLGO_FILL_RESIZE;
		return rc;
	}
	void ** p = h->handles;
	for (size_t i = 0; i < h->subscriptions; i++, p++) {
		rc = rcl_wait_set_add_subscription(ws, (const rcl_subscription_t *)*p, NULL);
		if (rc != RCL_RET_OK) {
			*failed = RCLGO_FILL_SUBSCRIPTION;
			return rc;
		}
	}
	for (size_t i = 0; i < h->timers; i++, p++) {
		rc = rcl_wait_set_add_timer(ws, (const rcl_timer_t *)*p, NULL);
		if (rc != RCL_RET_OK) {
			*failed = RCLGO_FILL_TIMER;
			return rc;
		}
	}
	for (size_t i = 0; i < h->services; i++, p++) {
		rc = rcl_wait_set_add_service(ws, (const rcl_service_t *)*p, NULL);
		if (rc != RCL_RET_OK) {
			*failed = RCLGO_FILL_SERVICE;
			return rc;
		}
	}
	for (size_t i = 0; i < h->clients; i++, p++) {
		rc = rcl_wait_set_add_client(ws, (const rcl_client_t *)*p, NULL);
		if (rc != RCL_RET_OK) {
			*failed = RCLGO_FILL_CLIENT;
			return rc;
		}
	}
	for (size_t i = 0; i < h->guard_conditions; i++, p++) {
		rc = rcl_wait_set_add_guard_condition(ws, (const rcl_guard_condition_t *)*p, NULL);
		if (rc != RCL_RET_OK) {
			*failed = RCLGO_FILL_GUARD_CONDITION;
			return rc;
		}
	}
	return RCL_RET_OK;
}
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

// handleCloses counts closes of entities whose handles WaitSets cache in C
// memory, so a WaitSet rebuilds its cache instead of passing a freed handle.
var handleCloses atomic.Uint64

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
	dirty   bool     // the entity slices changed since handles was built

	handles      C.rclgo_wait_handles_t
	handlesCap   int
	handlesClose uint64 // handleCloses when handles was built

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
		dirty:          true,
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
		w.dirty = true
		w.mu.Unlock()
		return
	}
	w.queue = append(w.queue, func() {
		w.mu.Lock()
		apply()
		w.dirty = true
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
	for _, res := range res.snapshot() {
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

// fillFailures names the rcl call that failed at each rclgo_wait_set_fill stage.
var fillFailures = [...]string{
	C.RCLGO_FILL_VALID:           "rcl_wait_set_is_valid()",
	C.RCLGO_FILL_CLEAR:           "rcl_wait_set_clear()",
	C.RCLGO_FILL_RESIZE:          "rcl_wait_set_resize()",
	C.RCLGO_FILL_SUBSCRIPTION:    "rcl_wait_set_add_subscription()",
	C.RCLGO_FILL_TIMER:           "rcl_wait_set_add_timer()",
	C.RCLGO_FILL_SERVICE:         "rcl_wait_set_add_service()",
	C.RCLGO_FILL_CLIENT:          "rcl_wait_set_add_client()",
	C.RCLGO_FILL_GUARD_CONDITION: "rcl_wait_set_add_guard_condition()",
}

// rebuildHandles copies the entity handles into C memory for
// rclgo_wait_set_fill, preserving the order dispatch indexes by.
func (w *WaitSet) rebuildHandles() {
	n := len(w.subscriptions) + len(w.timers) + len(w.services) + len(w.clients) + len(w.guardConditions)
	if n > w.handlesCap {
		C.free(unsafe.Pointer(w.handles.handles))
		w.handles.handles = (*unsafe.Pointer)(C.malloc(C.size_t(n) * C.size_t(unsafe.Sizeof(unsafe.Pointer(nil)))))
		w.handlesCap = n
	}
	handles := unsafe.Slice(w.handles.handles, n)[:0]
	for _, s := range w.subscriptions {
		handles = append(handles, unsafe.Pointer(s.rcl_subscription_t))
	}
	for _, t := range w.timers {
		handles = append(handles, unsafe.Pointer(t.rcl_timer_t))
	}
	for _, s := range w.services {
		handles = append(handles, unsafe.Pointer(s.rclService))
	}
	for _, c := range w.clients {
		handles = append(handles, unsafe.Pointer(c.rclClient))
	}
	for _, g := range w.guardConditions {
		handles = append(handles, unsafe.Pointer(g.rclGuardCondition))
	}
	w.handles.subscriptions = C.size_t(len(w.subscriptions))
	w.handles.timers = C.size_t(len(w.timers))
	w.handles.services = C.size_t(len(w.services))
	w.handles.clients = C.size_t(len(w.clients))
	w.handles.guard_conditions = C.size_t(len(w.guardConditions))
}

// initEntities refills the rcl wait set, which rcl_wait prunes on every call.
func (w *WaitSet) initEntities() error {
	if closes := handleCloses.Load(); w.dirty || closes != w.handlesClose {
		w.rebuildHandles()
		w.dirty = false
		w.handlesClose = closes
	}
	var failed C.int
	rc := C.rclgo_wait_set_fill(
		&w.rcl_wait_set_t,
		&w.handles,
		C.size_t(2*len(w.actionClients)),
		C.size_t(len(w.actionServers)),
		C.size_t(3*len(w.actionClients)),
		C.size_t(3*len(w.actionServers)),
		&failed,
	)
	if rc != C.RCL_RET_OK {
		return errorsCastC(rc, fmt.Sprintf("%s failed for wait set (%s)", fillFailures[failed], w.entityCounts()))
	}
	for _, server := range w.actionServers {
		rc = C.rcl_action_wait_set_add_action_server(&w.rcl_wait_set_t, &server.rclServer, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_action_server() failed for wait set (%s)", w.entityCounts()))
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
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_action_client() failed for wait set (%s)", w.entityCounts()))
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
	C.free(unsafe.Pointer(w.handles.handles))
	w.handles = C.rclgo_wait_handles_t{}
	w.handlesCap = 0
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
