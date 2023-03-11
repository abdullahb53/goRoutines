package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// When our worker is created by main function, we be able to collect our jobs.
	for i := range jobs {
		// fmt.Println("worker:", id, "Finished job..:", i)
		results <- i
	}
}

func main() {
	const Workers = 8
	const _JOBS = 10000000
	jobs := make(chan int, _JOBS)
	results := make(chan int, _JOBS)

	// Creating our workers that these are 5
	for w := 1; w <= Workers; w++ {
		go worker(w, jobs, results)
	}

	// Start timer after routines loaded.
	start := time.Now()

	// Sending jobs to 'jobs channel'.
	go func() {
		for j := 1; j <= _JOBS; j++ {
			jobs <- j + 331
		}
	}()

	// We are collecting our finished jobs with 'results channel'.

	i := 0
break_there:
	for {
		select {
		case <-results:
			i++

			if i >= _JOBS {
				break break_there
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}
