package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("../tmp/data/defer_demo.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("Creating file...")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("Writing some data into file...")
	fmt.Fprintln(f, "some data here")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file...")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
