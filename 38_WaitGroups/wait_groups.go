package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	// before 1.22
	/*
		for i := 1; i <= 5; i++ {
			wg.Add(1)

			go func(v int) {
				defer wg.Done()
				worker(v)
			}(i)
		}
	*/

	// After 1.22
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			worker(i)
			defer wg.Done()
		}()
	}

	wg.Wait()

}
