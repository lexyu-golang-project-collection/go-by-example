package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			time.Sleep(time.Second * 2)
			if more {
				fmt.Println("Received Job", j)
			} else {
				fmt.Println("Received All Job")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 5; i++ {
		j := "job-" + strconv.Itoa(i)
		jobs <- j
		fmt.Println("Sent Job", j)
		time.Sleep(time.Second)
	}
	close(jobs)
	fmt.Println("Sent All Job Done")

	<-done

	_, ok := <-jobs
	fmt.Println("Receive more jobs:", ok)
}
