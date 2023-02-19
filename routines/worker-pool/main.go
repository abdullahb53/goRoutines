package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// When our worker is created by main function, we be able to collect our jobs.
	for i := range jobs {
		fmt.Println("worker:", id, "Started job!:", i)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, "Finished job..:", i)
		results <- i * 3
	}
}

func main() {
	const numberJobs = 5
	jobs := make(chan int, numberJobs)
	results := make(chan int, numberJobs)

	// Creating our workers that these are 5
	// 'const numberJobs = 5'
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Sending jobs to 'jobs channel'.
	for j := 1; j <= numberJobs; j++ {
		jobs <- j + 3312
	}

	close(jobs)

	// We are collecting our finished jobs with 'results channel'.
	for a := 1; a <= numberJobs; a++ {
		<-results
	}
}
