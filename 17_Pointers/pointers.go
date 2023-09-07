package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
	fmt.Println("ival address:", &ival)
}

func zeroptr(iptr *int) {
	*iptr = 12345
	fmt.Println("iptr address:", &iptr)
}

func main() {
	i := 1
	value := &i
	a := *value
	fmt.Printf("initial: i=%d, a=%d\n", i, a)
	fmt.Printf("initial: i=%p, a=%p\n", &i, &a)

	zeroval(a)
	fmt.Println("zeroval:", a)

	zeroptr(&a)
	*value = 11111
	fmt.Printf("zeroptr: i=%d, a=%d\n", i, a)

	fmt.Println("pointer:", &a)
}
