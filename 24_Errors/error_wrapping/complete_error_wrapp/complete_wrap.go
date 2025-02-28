package main

import (
	"errors"
	"fmt"
)

func DoSomething() error {
	return errors.New("something didn't work")
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}

func main() {
	fmt.Println("err = ", DoSomething())

	_, e := Divide(1, 0)
	fmt.Println(e)
}
