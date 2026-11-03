package rclgo

import "unsafe"

var Testing_errorsCastC = errorsCastC

func (w *WaitSet) Testing_initEntities() error { return w.initEntities() }

// Testing_cachedHandles returns the entity handles initEntities last passed
// to rcl, in fill order.
func (w *WaitSet) Testing_cachedHandles() []unsafe.Pointer {
	n := w.handles.subscriptions + w.handles.timers + w.handles.services +
		w.handles.clients + w.handles.guard_conditions
	return append([]unsafe.Pointer(nil), unsafe.Slice(w.handles.handles, n)...)
}
