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

// Example
type Employee struct {
	Name   string
	Salary int
}

var employees = []Employee{
	{"A", 10},
	{"Z", 11},
	{"skip", 10000},
}

// func EmployeeIter(es []Employee) func(func(int, Employee) bool)
func EmployeeIter(es []Employee) iter.Seq2[int, Employee] {
	return func(yield func(int, Employee) bool) {
		for i := 0; i <= len(es)-1; i++ {
			if !yield(i, es[i]) {
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

	iter := EmployeeIter(employees)

	iter(func(i int, e Employee) bool {
		if e.Name != "skip" {
			fmt.Println("idx:", i, "name:", e.Name, "salary:", e.Salary)
		}
		return true
	})

	for i, e := range Iterate(employees) {
		fmt.Printf("%d: %+v\n", i, e)
	}

	for i, val := range Iterate([]int{1, 2, 3, 4, 5}) {
		fmt.Printf("%d: %+v\t", i, val)
	}
}
