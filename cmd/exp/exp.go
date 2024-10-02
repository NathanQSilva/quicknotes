package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("teste").Parse("<h1>Olá {{ . }}</h1>")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, "Nathan")
	if err != nil {
		panic(err)
	}
}
