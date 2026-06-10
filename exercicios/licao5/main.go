package main

import (
	"errors"
	"fmt"
)

type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {

	produtos := []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca", Preco: 3.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guarana", Preco: 3.99, Quantidade: 31},
	}

	p, err := buscProd(produtos, "003")
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("achei:", p.Nome, "R", p.Preco)
	}

	p, err = buscProd(produtos, "999")
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("achei:", p.Nome, "R", p.Preco)
	}

	fmt.Println(diviNumb(10, 5))
}

func diviNumb(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("deu ruim")
	} else {
		q := a / b
		return q, nil
	}

}

func buscProd(produtos []Produto, ccodigo string) (Produto, error) {

	for _, p := range produtos {
		if p.Codigo == ccodigo {
			return p, nil
		}
	}

	return Produto{}, fmt.Errorf("o produto com o codigo %s nao foi encontrado", ccodigo)

}
