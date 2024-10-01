package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header()["Date"] = nil // Suprimir header default

	fmt.Println(w, `{"id": 1}`)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.
		fmt.Println(w, "Visualização de nota")
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	// Rejeitar requisição
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(w, "Criação de nota")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
