package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:consys@tcp(localhost:3306)/consys"

type Produto struct {
	Codigo string  `json:"codigo"`
	Nome   string  `json:"nome"`
	Preco  float64 `json:"preco"`
}

func main() {
	http.HandleFunc("/api/produtos", produtos)
	log.Println("Conectou em: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func produtos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarProdutos(w, r)
	case http.MethodPost:
		inserirProduto(w, r)
	case http.MethodDelete:
		excluirProduto(w, r)
	// case http.MethodPut:
	// 	atualizarProduto(w, r)
	default:
		http.Error(w, "metodo nao suportado", http.StatusMethodNotAllowed)
	}
}

func listarProdutos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dsn)
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

	produtos := []Produto{}
	for rows.Next() {
		var p Produto
		if err := rows.Scan(&p.Codigo, &p.Nome, &p.Preco); err != nil {
			http.Error(w, "erro ao ler produto", http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	escreveJSON(w, http.StatusOK, produtos)
}

func inserirProduto(w http.ResponseWriter, r *http.Request) {
	var p Produto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		escreveJSON(w, http.StatusBadRequest, map[string]any{"erro": "JSON invalido"})
		return
	}

	p.Codigo = strings.TrimSpace(p.Codigo)
	p.Nome = strings.TrimSpace(p.Nome)
	if p.Codigo == "" || p.Nome == "" {
		escreveJSON(w, http.StatusBadRequest, map[string]any{"erro": "codigo e nome sao obrigatorios"})
		return
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao abrir o banco"})
		return
	}
	defer db.Close()

	var existe int
	err = db.QueryRow("SELECT COUNT(*) FROM alqui WHERE ccodigo = ? AND sr_deleted <> 'T'", p.Codigo).Scan(&existe)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao verificar produto"})
		return
	}
	if existe > 0 {
		escreveJSON(w, http.StatusConflict, map[string]any{"erro": "ja existe produto com este codigo"})
		return
	}

	res, err := db.Exec("INSERT INTO alqui (ccodigo, cdesc, cvenda, sr_deleted) VALUES (?, ?, ?, '')",
		p.Codigo, p.Nome, p.Preco)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao inserir: " + err.Error()})
		return
	}

	recno, _ := res.LastInsertId()
	escreveJSON(w, http.StatusCreated, map[string]any{
		"ok":     true,
		"recno":  recno,
		"codigo": p.Codigo,
		"nome":   p.Nome,
		"preco":  p.Preco,
	})
}

func excluirProduto(w http.ResponseWriter, r *http.Request) {

	codigo := strings.TrimSpace(r.URL.Query().Get("codigo"))
	if codigo == "" {
		var p Produto
		if err := json.NewDecoder(r.Body).Decode(&p); err == nil {
			codigo = strings.TrimSpace(p.Codigo)
		}
	}
	if codigo == "" {
		escreveJSON(w, http.StatusBadRequest, map[string]any{"erro": "codigo e obrigatorio"})
		return
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao abrir o banco"})
		return
	}
	defer db.Close()

	res, err := db.Exec("UPDATE alqui SET sr_deleted = 'T' WHERE ccodigo = ? AND sr_deleted <> 'T'", codigo)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao excluir: " + err.Error()})
		return
	}

	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		escreveJSON(w, http.StatusNotFound, map[string]any{"erro": "produto nao encontrado"})
		return
	}

	escreveJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"codigo": codigo,
	})
}

func atualizarProduto(w http.ResponseWriter, r *http.Request) {
	var p Produto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		escreveJSON(w, http.StatusBadRequest, map[string]any{"erro": "JSON invalido"})
		return
	}

	p.Codigo = strings.TrimSpace(p.Codigo)
	p.Nome = strings.TrimSpace(p.Nome)
	if p.Codigo == "" || p.Nome == "" {
		escreveJSON(w, http.StatusBadRequest, map[string]any{"erro": "codigo e nome sao obrigatorios"})
		return
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao abrir o banco"})
		return
	}
	defer db.Close()

	res, err := db.Exec("UPDATE alqui SET cdesc = ?, cvenda = ? WHERE ccodigo = ? AND sr_deleted <> 'T'",
		p.Nome, p.Preco, p.Codigo)
	if err != nil {
		escreveJSON(w, http.StatusInternalServerError, map[string]any{"erro": "erro ao atualizar: " + err.Error()})
		return
	}

	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		escreveJSON(w, http.StatusNotFound, map[string]any{"erro": "produto nao encontrado"})
		return
	}

	escreveJSON(w, http.StatusOK, map[string]any{
		"ok":     true,
		"codigo": p.Codigo,
		"nome":   p.Nome,
		"preco":  p.Preco,
	})
}

func escreveJSON(w http.ResponseWriter, status int, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "erro ao gerar JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
