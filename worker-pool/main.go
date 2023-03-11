package main

import (
	"fmt"
	"time"
)

const (
	workers = 8
	jobs    = 10000000
)

func worker(id int, jobsChannel <-chan int, resultsChannel chan<- int) {
	// When our worker is created by main function, we be able to collect our jobs.
	for i := range jobsChannel {
		// fmt.Println("worker:", id, "Finished job..:", i)
		resultsChannel <- i
	}
}

func main() {

	jobsChannel := make(chan int, jobs)
	resultsChannel := make(chan int, jobs)

	// Creating our workers that these are 5
	for w := 1; w <= workers; w++ {
		go worker(w, jobsChannel, resultsChannel)
	}

	// Start timer after routines loaded.
	start := time.Now()

	// Sending jobs to 'jobs channel'.
	go func() {
		for j := 1; j <= jobs; j++ {
			jobsChannel <- j + 331
		}
	}()

	// We are collecting our finished jobs with 'results channel'.

	i := 0
break_there:
	for {
		select {
		case <-resultsChannel:
			i++

			if i >= jobs {
				break break_there
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}
