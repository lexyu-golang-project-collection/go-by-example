package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)

	sum(9, 8, 7, 6, 5, 4, 3, 2, 1)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
