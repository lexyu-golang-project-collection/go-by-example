package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "deault-string", "a string")

	numbPtr := flag.Int("num", 100, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var str_var string
	flag.StringVar(&str_var, "svar", "default-svar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", str_var)
	fmt.Println("tail:", flag.Args())
}
