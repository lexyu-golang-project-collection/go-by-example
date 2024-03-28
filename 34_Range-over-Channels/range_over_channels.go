package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "ONE"
	queue <- "TWO"

	close(queue)

	for e := range queue {
		fmt.Println("Element:", e)
	}
}
