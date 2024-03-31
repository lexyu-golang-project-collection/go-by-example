package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(2 * time.Second)
		fmt.Println("worker", id, "finished  job", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 1; i <= 5; i++ {
		go worker(i, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= numJobs; r++ {
		rr := <-result
		fmt.Println("After Worker:", rr)
	}

}
