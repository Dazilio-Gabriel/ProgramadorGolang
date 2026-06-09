package main

import "fmt"
type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	produtos := listarProdutos()

	fmt.Println("\nLista de Produtos:")
	for i, p := range produtos {
		fmt.Printf("%d - %-10s | Preço: R$%.2f | Estoque: %d | Total: R$%.2f\n",
			i+1, p.Nome, p.Preco, p.Quantidade, p.ValorTotal())
	}

	totalGeral := valorTotalEstoque(produtos)
	fmt.Printf("\n💰 Valor Total em Estoque: R$%.2f\n", totalGeral)
}

func (p Produto) ValorTotal() float64 {
	return p.Preco * float64(p.Quantidade)
}

func listarProdutos() []Produto {
	return []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca Cola", Preco: 5.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guaraná", Preco: 3.50, Quantidade: 31},
	}
}

func valorTotalEstoque(produtos []Produto) float64 {
	total := 0.0
	for _, p := range produtos {
		total += p.ValorTotal()
	}
	return total
}