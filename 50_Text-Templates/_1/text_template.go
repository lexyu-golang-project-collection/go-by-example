package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1 = template.Must(t1.Parse("value is {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go", "Rust", "C++", "Java",
	})

}
