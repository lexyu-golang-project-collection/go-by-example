package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "Hello"
	}()

	time.Sleep(1 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	messages2 := make(chan string, 1)

	msg := "hi"
	select {
	case messages2 <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	messages3 := make(chan string)
	signals := make(chan bool)

	go func() {
		messages3 <- "World"
		signals <- true
	}()

	select {
	case msg := <-messages3:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
