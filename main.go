package main

import (
	"sync"
	"time"

	"github.com/speecan/quiz01/job"
)

var (
	maxWorkers   = 1000                          // maxWorkers means how many workers
	maxQueue     = 1000                          // maxQueue depends on your PCs memory
	numberOfJobs = 30000                         // numberOfJobs is all of jobs that have to work
	jobQueue     = make(chan *job.Job, maxQueue) // for queuing, use me efficiently
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < numberOfJobs; i++ {
		j := job.NewJob(i, 300*time.Millisecond)
		// SATISFY THIS SCOPE:
		{
			wg.Add(1) // add worker per job
			// number of queues = number of jobs
			// many jobs uses many memories (spending numberOfJobs x about 1 MB)
			go func(workerID int) {
				// worker
				j.Work(workerID) // work once
				wg.Done()        // notify waitgroup that i am done
			}(i)
		}
	}
	close(jobQueue)
	println("**** all jobs were queued ****")
	// wait for completing jobs
	wg.Wait()
}
