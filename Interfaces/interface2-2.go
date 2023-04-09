package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func DoSomething(v interface{}) {
	// ...
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	// cannot use (Cat literal) (value of type Cat) as Animal value in array or slice literal:
	// Cat does not implement Animal (method Speak has pointer receiver)
	// animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}

	// fix, new(Cat) / &Cat{}
	/*
		animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}
		for _, aanimal := range animals {
			fmt.Println(aanimal.Speak())
		}
	*/

	// everything in Go is passed by value

	animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
	for _, aanimal := range animals {
		fmt.Println(aanimal.Speak())
	}

}
