package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 50
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 60}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 9999
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
