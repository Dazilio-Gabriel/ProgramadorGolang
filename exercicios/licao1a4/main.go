package main

import "fmt"

type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	mostrarProduto()
	mostrarProdutosArray()
	fmt.Println(somaNumb(1, 4))
	fmt.Println(dividir(10, 5))
	fmt.Println(ehPar(4))
	tabuada(10)
	valoTota()
}

func somaNumb(a, b int) int {
	return a + b
}

func dividir(a, b int) (int, int) {
	return a / b, a % b
}

func ehPar(a int) bool {
	return a%2 == 0
}

func tabuada(a int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "- ", a*i)
	}
}

func valoTota() {
	produtos := []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca", Preco: 3.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guarana", Preco: 3.99, Quantidade: 31},
	}
	for i, p := range produtos {
		fmt.Println(i, "-", p.Nome, "valor do estoque", p.Preco*float64(p.Quantidade))
	}
}

func mostrarProduto() {
	var p1 Produto
	p1.Codigo = "001"
	p1.Nome = "Coca Cola"
	p1.Preco = 5.99
	p1.Quantidade = 4

	p2 := Produto{Codigo: "002", Nome: "Pepsi", Preco: 3.99, Quantidade: 5}

	fmt.Println("produto 1 =", p1)
	fmt.Println("produto 2 =", p2)
}

func mostrarProdutosArray() {
	produtos := []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca", Preco: 3.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guarana", Preco: 3.99, Quantidade: 31},
	}
	for _, p := range produtos {
		fmt.Println(p.Nome, p.Preco, p.Quantidade, p.Codigo)
	}
}
