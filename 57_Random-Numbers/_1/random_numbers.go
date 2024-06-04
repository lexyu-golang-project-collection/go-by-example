package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
}
