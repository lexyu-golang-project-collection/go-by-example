package main

import (
	"fmt"
	"io"
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

	o3, err := f.Seek(6, 0)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)

	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = f.Seek(0, 0)
	check(err)

	f.Close()
}
