package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Handler")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")

	mux := http.NewServeMux()
	// Rota coring / ou /hello/
	// Rota estática /hello
	mux.HandleFunc("/", homeHandler)

	http.ListenAndServe(":5000", mux)
}
