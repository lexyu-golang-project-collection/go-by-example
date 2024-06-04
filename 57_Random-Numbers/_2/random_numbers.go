package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+10, ",")
	fmt.Print(rand.Float64()*5 + 5)
}
