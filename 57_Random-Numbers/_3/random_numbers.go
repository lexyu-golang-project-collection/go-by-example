package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	seed2 := rand.NewPCG(720, 1024)
	r2 := rand.New(seed2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
}
