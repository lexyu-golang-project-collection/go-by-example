package main

import "fmt"

var pl = fmt.Println

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	pl(b, c)

	var d = true
	pl(d)

	var e int
	pl(e)

	f := "apple"
	pl(f)
}
