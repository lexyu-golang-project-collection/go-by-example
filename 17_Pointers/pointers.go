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
	fmt.Printf("initial: i=%d, a=%d, value=%d\n", i, a, value)
	fmt.Printf("initial: i=%p, a=%p, value=%p\n", &i, &a, &value)

	zeroval(a)
	fmt.Println("after zeroval a=", a)
	fmt.Printf("after zeroval: i=%p, a=%p\n", &i, &a)

	zeroptr(&a)
	fmt.Printf("after zeroptr: i=%d, a=%d\n", i, a)
	fmt.Printf("after zeroptr: i=%p, a=%p\n", &i, &a)

}
