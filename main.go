package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")

	handler := MyHandler{}

	http.ListenAndServe(":5000", handler)
}
