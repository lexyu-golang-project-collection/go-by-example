package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func main() {
	people := []*Person{
		&Person{name: "Jax", age: 37},
		&Person{name: "TJ", age: 25},
		&Person{name: "Alex", age: 100},
	}

	ageCmp := func(a, b *Person) int {
		return -cmp.Compare(a.age, b.age)
	}

	slices.SortFunc(people, ageCmp)
	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}
}
