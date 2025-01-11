package scheduler

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
)

// TestNewCronScheduler tests the creation of a new CronScheduler instance.
func TestNewCronScheduler(t *testing.T) {
	cs := NewCronScheduler()
	assert.NotNil(t, cs)
	assert.NotNil(t, cs.(*CronScheduler).CronInstance)
}

// TestScheduleTask tests scheduling a recurring task with a cron expression.
func TestScheduleTask(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	task := func() {}
	schedule := "0 0 * * *"
	id, err := cs.ScheduleTask(schedule, task)

	assert.NoError(t, err)
	assert.NotEqual(t, id, cron.EntryID(0))
}

// TestScheduleTaskLater tests scheduling a one-time task after a delay.
func TestScheduleTaskLater(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	task := func() {}
	delay := 1 * time.Second
	id := cs.ScheduleTaskLater(delay, task)

	assert.NotEqual(t, id, cron.EntryID(0))
}

// TestScheduleTasksLater tests scheduling multiple one-time tasks with different delays.
func TestScheduleTasksLater(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	delays := []time.Duration{1 * time.Second, 2 * time.Second}
	tasks := []func(){func() {}, func() {}}
	ids := cs.ScheduleTasksLater(delays, tasks)

	assert.Len(t, ids, 2)
	for _, id := range ids {
		assert.NotEqual(t, id, cron.EntryID(0))
	}
}

// TestScheduleRepeatingTask tests scheduling a repeating task with a fixed interval.
func TestScheduleRepeatingTask(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	task := func() {}
	interval := 2 * time.Second
	id := cs.ScheduleRepeatingTask(interval, task)

	assert.NotEqual(t, id, cron.EntryID(0))
}

// TestStop tests stopping the scheduler and its tasks.
func TestStop(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)
	cs.Start()

	cs.Stop()
	assert.Len(t, cs.CronInstance.Entries(), 0)
}

// TestStart tests starting the scheduler.
func TestStart(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)
	cs.Start()
	assert.NotNil(t, cs.CronInstance)
}

// TestScheduleTaskError tests that an error in scheduling a task returns an appropriate error.
func TestScheduleTaskError(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	// Using an invalid cron expression to induce an error
	task := func() {}
	schedule := "invalid cron expression"
	id, err := cs.ScheduleTask(schedule, task)

	// Assert that an error is returned and no valid ID is provided
	assert.Error(t, err)
	assert.Equal(t, id, cron.EntryID(0))
}

// TestScheduleTasksLaterPanics tests that ScheduleTasksLater panics when the number of delays and tasks don't match.
func TestScheduleTasksLaterPanics(t *testing.T) {
	cs := NewCronScheduler().(*CronScheduler)

	// Test when the lengths of delays and tasks arrays do not match
	delays := []time.Duration{1 * time.Second}
	tasks := []func(){func() {}, func() {}} // One task more than delays

	require.Panics(t, func() {
		cs.ScheduleTasksLater(delays, tasks)
	})
}
