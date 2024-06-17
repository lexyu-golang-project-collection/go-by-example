package main

import (
	"fmt"
	str "strings"
)

var pl = fmt.Println

func main() {
	pl("Contains: ", str.Contains("test", "e"))
	pl("Count:     ", str.Count("test", "t"))
	pl("HasPrefix: ", str.HasPrefix("test", "te"))
	pl("HasSuffix: ", str.HasSuffix("test", "st"))
	pl("Index:     ", str.Index("test", "e"))
	pl("Join:      ", str.Join([]string{"a", "b"}, "-"))
	pl("Repeat:    ", str.Repeat("a", 5))
	pl("Replace:   ", str.Replace("foo", "o", "0", -1))
	pl("Replace:   ", str.Replace("foo", "o", "0", 1))
	pl("Split:     ", str.Split("a-b-c-d-e", "-"))
	pl("ToLower:   ", str.ToLower("TEST"))
	pl("ToUpper:   ", str.ToUpper("test"))
}
