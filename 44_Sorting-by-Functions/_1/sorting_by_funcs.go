package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)
}
