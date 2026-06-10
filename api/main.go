package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/api/produtos", listarProdutosHandler)
	log.Println("servidor no ar → http://localhost:8080/api/produtos (Ctrl+C para parar)")
	http.ListenAndServe(":8080", nil)
}

func listarProdutosHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mostrarProdutosArray())
}

func mostrarProdutosArray() []Produto {
	return []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca", Preco: 3.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guarana", Preco: 3.99, Quantidade: 31},
	}
}
