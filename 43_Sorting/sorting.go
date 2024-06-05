package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "z", "a"}
	slices.Sort(strs)
	fmt.Println("strs:", strs)

	res := slices.IsSorted(strs)
	fmt.Println("Sorted:", res)

	slices.Reverse(strs)
	fmt.Println("strs:", strs)
	fmt.Println("Sorted:", res)

	nums := []int{100, 5, 2, 9}
	slices.Sort(nums)
	fmt.Println("nums:", nums)

	res = slices.IsSorted(nums)
	fmt.Println("Sorted:", res)
}
