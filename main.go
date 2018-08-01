package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/speecan/quiz01/job"
)

var (
	maxWorkers   = 10000   // maxWorkers means how many workers
	maxQueue     = 100000  // maxQueue depends on your PCs memory
	numberOfJobs = 1000000 // numberOfJobs is all of jobs that have to work
)

// Worker is so serious
// Please Satisfy ME!
type Worker struct {
	ID    int
	Queue chan job.Job
}

func main() {
	for i := 0; i < numberOfJobs; i++ {
		d := time.Duration(rand.Intn(1000)) * time.Millisecond
		j := job.NewJob(i, d)

		// SATISFYME:
		// Implement queue/worker flow
		// (this will be not going to panic since v1.10??)
		go j.Work(i)
		//
	}

	fmt.Println("all were queued")

	// wait forever
	http.ListenAndServe(":10000", nil)
}
