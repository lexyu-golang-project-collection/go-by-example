package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}

	// cannot use names (variable of type []string) as []interface{} value in argument to
	// PrintAll(names)

	// fix, convert the []string to an []interface{}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}
