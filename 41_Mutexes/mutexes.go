package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.counters[name]++
}

func main() {
	var wg sync.WaitGroup

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	doIncr := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)

	go doIncr("a", 1000)
	go doIncr("b", 1000)
	go doIncr("a", 1000)

	wg.Wait()

	fmt.Printf("%+v", c.counters)
}
