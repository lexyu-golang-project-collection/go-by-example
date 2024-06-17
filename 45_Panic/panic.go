package main

import "os"

func main() {
	panic("problem occur")

	_, err := os.Create("/some_file")
	if err != nil {
		panic(err)
	}
}
