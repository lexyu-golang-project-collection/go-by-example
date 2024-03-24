package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")

	time.Sleep(time.Second * 5)

	fmt.Println("Done!")

	done <- true
}

func main() {

	done := make(chan bool)

	go worker(done)

	/*
		If you removed the <-done line from this program, the program would exit before the worker even started.
	*/
	println(<-done)
}
