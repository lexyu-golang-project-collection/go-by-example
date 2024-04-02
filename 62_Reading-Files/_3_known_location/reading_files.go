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
	f, err := os.Open("tmp/data/rand.txt")
	check(err)

	o2, err := f.Seek(6, 0)
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)

	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	f.Close()
}
