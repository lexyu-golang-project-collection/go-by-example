package main

import (
	"os"
	"text/template"
)

func main() {
	t3 := Create("t3",
		"{{if . -}} yes!! {{else -}} no!! {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
}

func Create(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}
