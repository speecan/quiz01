package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/speecan/quiz01/job"
)

var (
	maxWorkers   = 1000  // maxWorkers means how many workers
	maxQueue     = 1000  // maxQueue depends on your PCs memory
	numberOfJobs = 30000 // numberOfJobs is all of jobs that have to work
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < numberOfJobs; i++ {
		j := job.NewJob(i, 300*time.Millisecond)
		// SATISFYME:
		wg.Add(1) // add worker per job
		// number of queues = number of jobs
		// many jobs uses many memories (spending numberOfJobs x about 1 MB)
		go func(workerID int) {
			j.Work(workerID)
			wg.Done()
		}(i)
		//
	}
	fmt.Println("all were queued")
	// wait for completing jobs
	wg.Wait()
}
