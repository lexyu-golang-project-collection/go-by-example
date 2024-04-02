package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("tmp/data/rand.txt")
	check(err)

	fmt.Println(data)
}
