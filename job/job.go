package job

import (
	"fmt"
	"time"
)

// NewJob creates and returns a new one
func NewJob(id int, duration time.Duration) *Job {
	return &Job{
		ID:       id,
		Duration: duration,
	}
}

// Job is a heavy work
type Job struct {
	ID       int
	Duration time.Duration
}

// Work a heavy job
func (j *Job) Work(workerID int) {
	fmt.Printf("worker %010d: job %010d started, working for %fs\n", workerID, j.ID, j.Duration.Seconds())
	time.Sleep(j.Duration)
	fmt.Printf("worker %010d: job %010d completed!\n", workerID, j.ID)
}
