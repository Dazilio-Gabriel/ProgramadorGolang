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
	http.ListenAndServe(":8080", nil)
}

func listaProdutosBancoDeDados(w http.ResponseWriter, r *http.Request) {
	var id int
	var codigo string
	var nome string
	var preco float64
	var quantidade int

	ctx := context.Background()

	db, err := sql.Open("mysql", "root:consys@tcp(localhost:3306)/apigo")
	if err != nil {
		log.Fatal("erro ao abrir o banco: ", err)
	}

	err = db.Ping()
	if err != nil {
		http.Error(w, "deu erro ao buscar produtos", 500)
		return
	}
	log.Println("conectado ao MySQL")

	rows, err := db.QueryContext(ctx, "SELECT * FROM PRODUTOS")
	if err != nil {
		log.Fatal("Select incorreto")
	}
	defer rows.Close()

	var vetor []Produto
	for rows.Next() {
		err = rows.Scan(&id, &codigo, &nome, &preco, &quantidade)
		if err != nil {
			log.Fatal(err)
		}
		p := Produto{Codigo: codigo, Nome: nome, Preco: preco, Quantidade: quantidade}
		vetor = append(vetor, p)
	}

	json.NewEncoder(w).Encode(vetor)
}
