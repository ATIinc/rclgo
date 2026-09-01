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

type WaitSet struct {
	rosID
	// WaitTimeout bounds each underlying rcl_wait call made by Run. Zero
	// means DefaultWaitTimeout; a negative value waits forever, which is the
	// historical behavior and leaves Run exposed to lost RMW wakeups.
	WaitTimeout     time.Duration
	Subscriptions   []*Subscription
	Timers          []*Timer
	Services        []*Service
	Clients         []*Client
	ActionClients   []*ActionClient
	ActionServers   []*ActionServer
	guardConditions []*guardCondition
	rcl_wait_set_t  C.rcl_wait_set_t
	cancelWait      *guardCondition
	context         *Context
	lostWakeups     atomic.Uint64
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
		Subscriptions:  []*Subscription{},
		Timers:         []*Timer{},
		Services:       []*Service{},
		Clients:        []*Client{},
		rcl_wait_set_t: C.rcl_get_zero_initialized_wait_set(),
	}
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
	ws.addGuardConditions(ws.cancelWait)
	c.addResource(ws)
	return ws, nil
}

// Context returns the context s belongs to.
func (w *WaitSet) Context() *Context {
	return w.context
}

func (w *WaitSet) AddSubscriptions(subs ...*Subscription) {
	w.Subscriptions = append(w.Subscriptions, subs...)
}

func (w *WaitSet) AddTimers(timers ...*Timer) {
	w.Timers = append(w.Timers, timers...)
}

func (w *WaitSet) AddServices(services ...*Service) {
	w.Services = append(w.Services, services...)
}

func (w *WaitSet) AddClients(clients ...*Client) {
	w.Clients = append(w.Clients, clients...)
}

func (w *WaitSet) AddActionServers(servers ...*ActionServer) {
	w.ActionServers = append(w.ActionServers, servers...)
}

func (w *WaitSet) AddActionClients(clients ...*ActionClient) {
	w.ActionClients = append(w.ActionClients, clients...)
}

func (w *WaitSet) addGuardConditions(guardConditions ...*guardCondition) {
	w.guardConditions = append(w.guardConditions, guardConditions...)
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
		len(w.Subscriptions), len(w.Timers), len(w.Services), len(w.Clients),
		len(w.ActionServers), len(w.ActionClients),
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
	for _, subscription := range w.Subscriptions {
		if subscription.waitable.reserve() {
			defer subscription.waitable.release()
		}
	}
	for _, timer := range w.Timers {
		if timer.waitable.reserve() {
			defer timer.waitable.release()
		}
	}
	for _, service := range w.Services {
		if service.waitable.reserve() {
			defer service.waitable.release()
		}
	}
	for _, client := range w.Clients {
		if client.waitable.reserve() {
			defer client.waitable.release()
		}
	}
	for _, actionClient := range w.ActionClients {
		if actionClient.waitable.reserve() {
			defer actionClient.waitable.release()
		}
	}
	for _, actionServer := range w.ActionServers {
		if actionServer.waitable.reserve() {
			defer actionServer.waitable.release()
		}
	}
	for _, guardCondition := range w.guardConditions {
		if guardCondition.waitable.reserve() {
			defer guardCondition.waitable.release()
		}
	}
	if ctx == nil {
		return errors.New("context must not be nil")
	}
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
		for i := range w.guardConditions {
			if guardConditions[i] == w.cancelWait.rclGuardCondition {
				return ctx.Err()
			}
		}

		if recoveredLostWakeup(prevTimedOut, elapsed) {
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
		timers := unsafe.Slice(w.rcl_wait_set_t.timers, len(w.Timers))
		for i, t := range w.Timers {
			if timers[i] != nil {
				t.Reset() //nolint:errcheck
				t.Callback(t)
			}
		}
		subs := unsafe.Slice(w.rcl_wait_set_t.subscriptions, len(w.Subscriptions))
		for i, s := range w.Subscriptions {
			if subs[i] != nil {
				s.Callback(s)
			}
		}
		svcs := unsafe.Slice(w.rcl_wait_set_t.services, len(w.Services))
		for i, s := range w.Services {
			if svcs[i] != nil {
				s.handleRequest()
			}
		}
		clients := unsafe.Slice(w.rcl_wait_set_t.clients, len(w.Clients))
		for i, c := range w.Clients {
			if clients[i] != nil {
				c.sender.HandleResponse()
			}
		}
		for _, s := range w.ActionServers {
			s.handleReadyEntities(ctx, w)
		}
		for _, c := range w.ActionClients {
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
		C.size_t(len(w.Subscriptions)+2*len(w.ActionClients)),
		C.size_t(len(w.guardConditions)),
		C.size_t(len(w.Timers)+len(w.ActionServers)),
		C.size_t(len(w.Clients)+3*len(w.ActionClients)),
		C.size_t(len(w.Services)+3*len(w.ActionServers)),
		w.rcl_wait_set_t.size_of_events,
	)
	if rc != C.RCL_RET_OK {
		return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_resize() failed for wait_set='%v'", w))
	}
	for _, sub := range w.Subscriptions {
		rc = C.rcl_wait_set_add_subscription(&w.rcl_wait_set_t, sub.rcl_subscription_t, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_subscription() failed for wait_set='%v'", w))
		}
	}
	for _, timer := range w.Timers {
		rc = C.rcl_wait_set_add_timer(&w.rcl_wait_set_t, timer.rcl_timer_t, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_timer() failed for wait_set='%v'", w))
		}
	}
	for _, service := range w.Services {
		rc = C.rcl_wait_set_add_service(&w.rcl_wait_set_t, service.rclService, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_service() failed for wait_set='%v'", w))
		}
	}
	for _, client := range w.Clients {
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
	for _, server := range w.ActionServers {
		rc = C.rcl_action_wait_set_add_action_server(&w.rcl_wait_set_t, &server.rclServer, nil)
		if rc != C.RCL_RET_OK {
			return errorsCastC(rc, fmt.Sprintf("rcl_wait_set_add_action_server() failed for wait_set='%v'", w))
		}
	}
	for _, client := range w.ActionClients {
		rc = C.rcl_action_wait_set_add_action_client(&w.rcl_wait_set_t, &client.rclClient, nil, nil)
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
	w.context.removeResource(w)
	w.context = nil
	rc := C.rcl_wait_set_fini(&w.rcl_wait_set_t)
	if rc != C.RCL_RET_OK {
		err = errors.Join(err, errorsCast(rc))
	}
	var closeError closeError
	cancelWaitErr := w.cancelWait.Close()
	if cancelWaitErr != nil && !errors.As(cancelWaitErr, &closeError) {
		err = errors.Join(err, cancelWaitErr)
	}
	return err
}
