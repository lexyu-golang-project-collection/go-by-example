package main

import (
	"errors"
	"fmt"
)

// Defining Sentinel Errors
var ErrDivideByZero = errors.New("divide by zero")

func Divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Constructing Errors
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

	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
