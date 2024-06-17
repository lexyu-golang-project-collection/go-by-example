package main

import "fmt"

func panicOccur() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panicOccur(). Error:\n", r)
		}
	}()

	panic("This is a Problem!")
}
func main() {

	panicOccur()

	fmt.Println("End - After panicOccur()")
}
