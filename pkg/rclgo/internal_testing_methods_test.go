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

func (s *ActionServer) Testing_takeGoalRequest() error {
	return s.takeGoalRequest(newEmptyGoal(s, func() {}))
}

// Testing_setResultOnEmptyGoal runs GoalHandle.setResult for a goal s never
// accepted, which drives rcl_action_notify_goal_done.
func (s *ActionServer) Testing_setResultOnEmptyGoal() {
	newEmptyGoal(s, func() {}).setResult(nil)
}
