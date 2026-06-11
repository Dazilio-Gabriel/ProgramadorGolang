package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {

	http.HandleFunc("/api/produtos", listaProdutosBancoDeDados)
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}

func listaProdutosBancoDeDados(w http.ResponseWriter, r *http.Request) {
	var codigo string
	var nome string
	var preco float64

	ctx := context.Background()

	db, err := sql.Open("mysql", "root:consys@tcp(localhost:3306)/consys")
	if err != nil {
		log.Fatal("erro ao abrir o banco: ", err)
	}

	err = db.Ping()
	if err != nil {
		http.Error(w, "deu erro ao buscar produtos", 500)
		return
	}
	log.Println("conectado ao MySQL")

	rows, err := db.QueryContext(ctx, "SELECT sr_recno,ccodigo,cdesc,cvenda FROM alqui")
	if err != nil {
		log.Fatal("Select incorreto")
	}
	defer rows.Close()

	var vetor []Produto
	for rows.Next() {
		err = rows.Scan(&codigo, &nome, &preco)
		if err != nil {
			log.Fatal(err)
		}
		p := Produto{Codigo: codigo, Nome: nome, Preco: preco}
		vetor = append(vetor, p)
	}

	json.NewEncoder(w).Encode(vetor)
}
