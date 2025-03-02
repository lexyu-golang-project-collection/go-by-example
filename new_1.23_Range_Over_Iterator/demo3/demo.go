package main

import (
	"fmt"
	"iter"
	"time"
)

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

func EmployeeIter(es []Employee) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i <= len(es)-1; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func EmployeeIter2(es []Employee) iter.Seq2[int, Employee] {
	return func(yield func(int, Employee) bool) {
		for i := 0; i <= len(es)-1; i++ {
			if !yield(i, es[i]) {
				return
			}
		}
	}
}

func Iterate[E any](e []E) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i <= len(e)-1; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Iterate2[E any](e []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i <= len(e)-1; i++ {
			time.Sleep(1 * time.Second)
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func IntSeq(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func IntSeq2(start, end int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		index := 0
		for i := start; i <= end; i++ {
			if !yield(index, i) {
				return
			}
			index++
		}
	}
}

func main() {
	IntIterDemos()
}

func EmployeeIterDemos() {
	empIter2 := EmployeeIter2(employees)

	empIter2(func(i int, e Employee) bool {
		if e.Name != "skip" {
			fmt.Println("idx:", i, "name:", e.Name, "salary:", e.Salary)
		}
		return true
	})

	for i, e := range Iterate2(employees) {
		fmt.Printf("%d: %+v\n", i, e)
	}

	fmt.Println("\n=== 使用 Pull2 方法 ===")

	fmt.Println("\n=== 使用 All 方法 ===")

	fmt.Println("\n=== 使用 Backward 方法 ===")
}

func IntIterDemos() {
	fmt.Println("\n=== 使用單值迭代器 ===")
	intSeq := IntSeq(1, 5)
	intSeq(func(val int) bool {
		if val%2 != 0 {
			fmt.Printf("value: %d\n", val)
		}
		return true
	})

	intSeq2 := IntSeq2(1, 5)
	intSeq2(func(idx, val int) bool {
		fmt.Printf("idx: %d, value: %d\n", idx, val)
		return true
	})

	fmt.Println("\n使用 for range:")
	for val := range intSeq {
		fmt.Printf("value: %d\n", val)
	}

	for idx, val := range intSeq2 {
		fmt.Printf("idx: %d, value: %d\n", idx, val)
	}

	fmt.Println("\n使用 Pull:")
	pull1, stop1 := iter.Pull(intSeq)
	defer stop1()
	for {
		val, ok := pull1()
		if !ok {
			fmt.Printf("not ok = %+v\n", ok)
			break
		}
		fmt.Printf("value: %d\n", val)
	}

	fmt.Println("\n使用 Pull2:")
	pull2, stop2 := iter.Pull2(intSeq2)
	defer stop2()
	// 1.
	for {
		idx, val, ok := pull2()
		if !ok {
			fmt.Printf("not ok = %+v\n", ok)
			break
		}
		fmt.Printf("idx: %d, value: %d\n", idx, val)
	}
	// 2.
	// for idx, val, ok := pull(); ok; idx, val, ok = pull() {
	// 	fmt.Printf("idx: %d, value: %d\n", idx, val)
	// }

}
