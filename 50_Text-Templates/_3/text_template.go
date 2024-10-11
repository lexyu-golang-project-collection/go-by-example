package main

import (
	"os"
	"text/template"
)

func main() {

	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{"Unknown Tester"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Unknown Tester"})
}

func Create(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}
