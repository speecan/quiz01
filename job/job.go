package job

import (
	"fmt"
	"time"
)

// NewJob creates and returns a new one
func NewJob(jobID int, duration time.Duration) *Job {
	return &Job{
		ID:       jobID,
		Duration: duration,
	}
}

// Job is a heavy work and size
type Job struct {
	ID        int
	Duration  time.Duration
	DummyData [1024 * 1024]byte // Job uses 1 MB memory until working
}

// Work a heavy job
func (j *Job) Work(workerID int) {
	fmt.Printf("worker %06d: job %06d started, working for %fs\n", workerID, j.ID, j.Duration.Seconds())
	time.Sleep(j.Duration)
	fmt.Printf("worker %06d: job %06d completed!\n", workerID, j.ID)
}
