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
	maxQueue     = 1000000 // maxQueue depends on your PCs memory
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

		// FIXME:
		// its going to CRASH !!!!!!!!!!
		go j.Work(0)
		//
	}

	fmt.Println("all were queued")

	// wait forever
	http.ListenAndServe(":10000", nil)
}
