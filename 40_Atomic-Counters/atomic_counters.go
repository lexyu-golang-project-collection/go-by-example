package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {

		go func() {
			for i := 0; i < 1000; i++ {

			}
		}()
	}

	fmt.Println("ops:", ops.Load())
}
