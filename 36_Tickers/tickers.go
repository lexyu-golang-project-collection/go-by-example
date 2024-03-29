package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case res := <-done:
				fmt.Println(res)
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(4 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
