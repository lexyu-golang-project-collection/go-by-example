package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	seed3 := rand.NewPCG(720, 1024)
	r3 := rand.New(seed3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
}
