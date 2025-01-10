package scheduler

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Scheduler defines an interface for scheduling tasks.
type Scheduler interface {
	// ScheduleTask schedules a recurring task using a cron expression.
	ScheduleTask(schedule string, task func()) (cron.EntryID, error)

	// ScheduleTaskLater schedules a one-time task after a delay.
	ScheduleTaskLater(delay time.Duration, task func()) cron.EntryID

	// ScheduleTasksLater schedules multiple one-time tasks with different delays.
	ScheduleTasksLater(delays []time.Duration, tasks []func()) []cron.EntryID

	// ScheduleRepeatingTask schedules a task to run repeatedly with a fixed interval.
	ScheduleRepeatingTask(interval time.Duration, task func()) cron.EntryID

	// Stop stops the scheduler and all its tasks.
	Stop()

	// Start starts the scheduler.
	Start()
}

// CronScheduler implements the Scheduler interface using robfig/cron.
type CronScheduler struct {
	cronInstance *cron.Cron
}

// NewCronScheduler creates a new CronScheduler instance.
func NewCronScheduler() Scheduler {
	return &CronScheduler{
		cronInstance: cron.New(),
	}
}

// ScheduleTask schedules a recurring task using a cron expression.
func (cs *CronScheduler) ScheduleTask(schedule string, task func()) (cron.EntryID, error) {
	id, err := cs.cronInstance.AddFunc(schedule, task)
	if err != nil {
		return 0, fmt.Errorf("failed to schedule task: %w", err)
	}
	return id, nil
}

// ScheduleTaskLater schedules a one-time task after a delay.
func (cs *CronScheduler) ScheduleTaskLater(delay time.Duration, task func()) cron.EntryID {
	var id cron.EntryID
	id = cs.cronInstance.Schedule(
		cron.Every(delay),
		cron.FuncJob(func() {
			task()
			cs.cronInstance.Remove(id)
		}),
	)
	return id
}

// ScheduleTasksLater schedules multiple one-time tasks with different delays.
func (cs *CronScheduler) ScheduleTasksLater(delays []time.Duration, tasks []func()) []cron.EntryID {
	if len(delays) != len(tasks) {
		panic("the number of delays and tasks must match")
	}

	var ids []cron.EntryID
	for i, delay := range delays {
		task := tasks[i]
		id := cs.ScheduleTaskLater(delay, task)
		ids = append(ids, id)
	}
	return ids
}

// ScheduleRepeatingTask schedules a task to run repeatedly with a fixed interval.
func (cs *CronScheduler) ScheduleRepeatingTask(interval time.Duration, task func()) cron.EntryID {
	id := cs.cronInstance.Schedule(cron.Every(interval), cron.FuncJob(task))
	return id
}

// Stop stops the scheduler and all its tasks.
func (cs *CronScheduler) Stop() {
	cs.cronInstance.Stop()
}

// Start starts the scheduler.
func (cs *CronScheduler) Start() {
	cs.cronInstance.Start()
}
