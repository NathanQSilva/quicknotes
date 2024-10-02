package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Nome  string
	Idade int
}

func main() {
	t, err := template.ParseFiles(
		"layout1.html",
		`layout2.html`,
		"home.html",
		"footer.html",
		"header.html",
	)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "layout1.html")
	if err != nil {
		panic(err)
	}
}
