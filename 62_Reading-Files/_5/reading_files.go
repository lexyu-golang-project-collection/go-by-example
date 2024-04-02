package main

import (
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

	_, err = f.Seek(0, 0)
	check(err)

	f.Close()
}
