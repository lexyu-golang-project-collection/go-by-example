package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type goroutineIDKey struct{}

var goroutineIDCounter uint64

func nextGoroutineID() uint64 {
	return atomic.AddUint64(&goroutineIDCounter, 1)
}

func withGoroutineID(ctx context.Context) context.Context {
	return context.WithValue(ctx, goroutineIDKey{}, nextGoroutineID())
}

func goroutineID(ctx context.Context) uint64 {
	id, _ := ctx.Value(goroutineIDKey{}).(uint64)
	return id
}

func worker(ctx context.Context, wg *sync.WaitGroup, counter *Counter) {
	defer wg.Done()
	id := goroutineID(ctx)
	value := counter.Increment()
	fmt.Printf("Goroutine %d is working, coutner = %+v\n", id, value)
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ctx := withGoroutineID(context.Background())
			worker(ctx, &wg, counter)
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
