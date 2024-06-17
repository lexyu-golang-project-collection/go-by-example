package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Stateful Goroutine
	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
				// fmt.Printf("state[read.key] = %[1]d, read.key = %[2]d\n", state[read.key], read.key)
			case write := <-writes:
				state[write.key] = write.val
				// fmt.Printf("state[write.key] = %[1]d, write.key = %[2]d, write.val = %[3]d\n", state[write.key], write.key, write.val)
				write.resp <- true
			}
		}
	}()

	// Read Goroutines
	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				// fmt.Printf("Read Goroutines - key = %[1]d\n", read.key)
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Write Goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(10),
					resp: make(chan bool),
				}
				writes <- write
				// fmt.Printf("Write Goroutines - key = %[1]d, value = %[2]d\n", write.key, write.val)
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(2 * time.Second)

	readOpsResult := atomic.LoadUint64(&readOps)
	writeOpsResult := atomic.LoadUint64(&writeOps)
	fmt.Println("Read Operations Result:", readOpsResult)
	fmt.Println("Write Operations Result:", writeOpsResult)
}
