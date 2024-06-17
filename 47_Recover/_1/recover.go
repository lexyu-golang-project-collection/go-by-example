package main

import "fmt"

func panicOccur() {
	panic("This is a Problem!")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panicOccur(). Error:\n", r)
		}
	}()

	panicOccur()

	fmt.Println("End - After panicOccur()") // won't occur
}
