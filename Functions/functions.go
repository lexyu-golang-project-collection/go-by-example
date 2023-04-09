package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(3, 4, 5)
	fmt.Println("3 + 4 + 5 = ", res)
}
