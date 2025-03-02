package main

import (
	"fmt"
	"iter"
	"time"
)

func Countdown(v int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := v; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// improve syntax
func CountdownR(v int) iter.Seq[int] {
	return nil
}

// Generics
// func Iterate[E any](e []E) func(func(int, E) bool)
func Iterate[E any](e []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i <= len(e)-1; i++ {
			time.Sleep(1 * time.Second)
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func main() {
	for x := range Countdown(10) {
		fmt.Printf("%d,\t", x)
	}

	fmt.Println()

	for i, val := range Iterate([]int{1, 2, 3, 4, 5}) {
		fmt.Printf("%d: %+v\t", i, val)
	}
}
