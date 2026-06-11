package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Codigo string  `json:"codigo"`
	Nome   string  `json:"nome"`
	Preco  float64 `json:"preco"`
}

func main() {
	http.HandleFunc("/api/produtos", listarProdutos)
	log.Println("servidor no ar -> http://localhost:8080/api/produtos")
	http.ListenAndServe(":8080", nil)
}

func listarProdutos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:consys@tcp(localhost:3306)/consys")
	if err != nil {
		http.Error(w, "erro ao abrir o banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT ccodigo, cdesc, cvenda FROM alqui where sr_deleted <> 'T'")
	if err != nil {
		http.Error(w, "erro ao consultar produtos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var p Produto
		if err := rows.Scan(&p.Codigo, &p.Nome, &p.Preco); err != nil {
			http.Error(w, "erro ao ler produto", http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	data, err := json.Marshal(produtos)
	if err != nil {
		http.Error(w, "erro ao gerar JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
}
