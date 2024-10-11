package main

import (
	"os"
	"text/template"
)

func main() {
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"A", "B", "C", "D",
	})
}

func Create(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}
