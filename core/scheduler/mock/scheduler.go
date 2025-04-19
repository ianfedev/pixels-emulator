package scheduler_test

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/mock"
)

// MockScheduler is a mock implementation of the Scheduler interface.
type MockScheduler struct {
	mock.Mock
}

// ScheduleTask mocks scheduling a recurring task using a cron expression.
func (m *MockScheduler) ScheduleTask(schedule string, task func()) (cron.EntryID, error) {
	args := m.Called(schedule, task)
	return args.Get(0).(cron.EntryID), args.Error(1)
}

// ScheduleTaskLater mocks scheduling a one-time task after a delay.
func (m *MockScheduler) ScheduleTaskLater(delay time.Duration, task func()) cron.EntryID {
	args := m.Called(delay, task)
	return args.Get(0).(cron.EntryID)
}

// ScheduleTasksLater mocks scheduling multiple one-time tasks with different delays.
func (m *MockScheduler) ScheduleTasksLater(delays []time.Duration, tasks []func()) []cron.EntryID {
	args := m.Called(delays, tasks)
	return args.Get(0).([]cron.EntryID)
}

// ScheduleRepeatingTask mocks scheduling a task to run repeatedly with a fixed interval.
func (m *MockScheduler) ScheduleRepeatingTask(interval time.Duration, task func()) cron.EntryID {
	args := m.Called(interval, task)
	return args.Get(0).(cron.EntryID)
}

// Stop mocks stopping the scheduler and all its tasks.
func (m *MockScheduler) Stop() {
	m.Called()
}

// Start mocks starting the scheduler.
func (m *MockScheduler) Start() {
	m.Called()
}

// Cancel mocks cancelling a scheduled task by its EntryID.
func (m *MockScheduler) Cancel(id cron.EntryID) {
	m.Called(id)
}
