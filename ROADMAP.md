# Roadmap — API de Produtos em Go (Consys)

Plano de aprendizado para sair do **Harbour/xHarbour** e construir uma API em **Go** + tela em **HTML/JS**, conectada ao MySQL (`check_fl02`).

Feito para quem **já sabe programar** (CRUD, SQL, validação, browse/formulário) — o foco é no que **muda** do Harbour pro Go.

---

## A arquitetura

No Consys/Harbour, o mesmo `.exe` desenha a tela **e** fala com o banco. No mundo web isso se parte em dois:

```
[ NAVEGADOR ]                       [ SEU .EXE EM GO ]              [ MySQL ]
HTML + JS  --- AJAX (fetch) --->    API REST (net/http)  ---SQL--->  check_fl02
  (a tela)     JSON vai e volta       (o "backend")                  (produtos)
```

- O **Go** faz o papel do `oBase:AbreTabe()/ValoCamp()`: abre conexão, roda SQL, devolve dados — mas em **JSON** em vez de desenhar `ConsBrow`.
- O **HTML/JS** faz o papel do `ConsForm`/`ConsBrow`: mostra a tabela e o formulário no navegador. O `fetch()` do JS é o "AJAX".

---

## Mapa mental: Harbour → Go

| No Harbour                            | No Go                                                    |
|---------------------------------------|---------------------------------------------------------|
| `LOCAL lcNome := Space(40)`           | `nome := ""` (o `:=` é igual! mas o tipo é fixo)         |
| Tipagem dinâmica                      | **Tipagem estática** — `string` é `string`, não vira número |
| Prefixo húngaro (`lc`, `ln`, `ld`)    | Nome curto camelCase, sem prefixo                       |
| `FUNCTION fCadsClie(...)`             | `func cadastraCliente(...)`                              |
| Classe `ConsForm():New()` + métodos   | `struct` + métodos (tem método, **não tem herança**)    |
| `aArraBrow` (array que cresce)        | `[]Produto` (slice — array dinâmico)                    |
| Hash / array associativo              | `map[string]int`                                        |
| `IF !Empty(...)` validação            | `if err != nil` — **erro é retornado, não exceção**     |
| `oBase:AbreTabe('PRODUTOS')`          | `database/sql` + driver MySQL                           |
| `xbuild` gera o `.exe` (linkando)     | `go build` gera **um .exe sozinho**, zero dependência   |

> Maior quebra de cabeça: **tratamento de erro**. Quase toda função que pode falhar devolve `(resultado, erro)`, e você checa `if err != nil`. Não tem `BEGIN SEQUENCE`.

---

## As fases

Cada fase tem **objetivo** e **"pronto quando"** (como saber que terminou). Não pule fases.

### Fase 0 — Ambiente  ✅ COMPLETA
- [x] Go instalado (`go version` → go1.26.0)
- [x] VS Code + extensão **"Go"** (da Google)
- [x] 3 comandos: `go mod init`, `go run .`, `go build`
- [x] `fmt.Println(...)` rodando com `go run .`

### Fase 1 — A linguagem, no terminal  🟡 *em andamento*
Conceitos:
- [x] variáveis, `:=` vs `=`, tipagem estática (Go barra erro de tipo na compilação)
- [x] público/privado pela **maiúscula** (e por que `main` é minúsculo)
- [x] `struct` — o "registro de produto"
- [x] `slice` (`[]Produto`) — a lista / o "browse"; índice começa em **0**
- [x] `for ... range` + o `_` (blank identifier) e a regra de variável não usada
- [x] métodos com *receiver* (ex.: `ValorTotal()`) — conceito visto
- [ ] **funções + retorno de erro** (`func ... (x, error)` + `if err != nil`)  ← próximo
- [ ] `map`, ponteiros (`*` / `&`) com calma
- [ ] (reforço) [Tour of Go](https://go.dev/tour) + [Go by Example](https://gobyexample.com)
- **Pronto quando:** criar `struct Produto` + `[]Produto` + `for range` → ✅ feito

### Fase 2 — Primeiro servidor HTTP (2 a 3 dias)
- [ ] Pacotes `net/http` (servidor) e `encoding/json` (devolver JSON)
- [ ] Rota `GET /api/produtos` devolvendo uma lista **fixa, hardcoded** em JSON
- [ ] Testar com navegador e com **Thunder Client** (extensão do VS Code)
- **Pronto quando:** abrir `http://localhost:8080/api/produtos` e ver o JSON

### Fase 3 — Conectar no MySQL e dar o SELECT (3 a 5 dias)  ⭐ *primeiro marco real*
- [ ] `database/sql` + driver `github.com/go-sql-driver/mysql` (`go get`)
- [ ] String de conexão: user `check`, senha `consys`, banco `check_fl02`, host `localhost`
- [ ] Descobrir a tabela de produtos real do Consys e suas colunas
- [ ] `GET /api/produtos` roda `SELECT` real e devolve JSON
- **Pronto quando:** a rota mostrar os produtos reais do banco

### Fase 4 — A tela: HTML/JS via AJAX (3 a 5 dias)
- [ ] Página HTML; no JS, `fetch('/api/produtos')` → montar uma `<table>`
- [ ] O Go também serve os arquivos estáticos (`http.FileServer`)
- **Pronto quando:** abrir a página e ver a tabela preenchida pelo banco

### Fase 5 — CRUD completo (1 a 2 semanas)  ⭐
- [ ] `POST /api/produtos` (criar), `PUT /api/produtos/{id}` (editar), `DELETE /api/produtos/{id}`
- [ ] **Prepared statements** (`?` nos parâmetros) — proteção contra SQL injection
- [ ] Validação no servidor (o que o `SayGetP` fazia com `{|x| !Empty(x)}`)
- [ ] Frontend: formulário de criar/editar + botão de excluir
- **Pronto quando:** cadastrar, alterar e apagar um produto pela tela e ver refletir no banco

### Fase 6 — Organizar o projeto (contínuo)
- [ ] Separar em pacotes: `handlers`, `repository`/`models`, `db`
- [ ] Ler config de arquivo/variável de ambiente (igual o `Consys_check.ini`), tirar senha do código
- [ ] Opcionais quando incomodar: router `chi`, `sqlx`, middleware de log

### Fase 7 — Próximos passos
- [ ] Login/autenticação
- [ ] Outros cadastros reaproveitando o padrão
- [ ] `go build` gerando o `.exe` final (binário único, sem DLL)

---

## Caminho crítico (resumo)
Instalar → linguagem no terminal → servidor que devolve JSON fixo → mesmo servidor lendo o MySQL → tela lendo via fetch → CRUD.

## Dados de conexão (do `Consys_check.ini`)
```
TIPOBASE = MYSQL
DTB      = check_fl02
HOST     = localhost
UID      = check
PWD      = consys
```

## Comandos Go do dia a dia
| Comando            | O que faz                                  |
|--------------------|--------------------------------------------|
| `go mod init nome` | Cria o `go.mod` (manifesto do projeto)     |
| `go run .`         | Compila e roda na hora (pra testar)        |
| `go build`         | Gera o `.exe`                              |
| `go get pacote`    | Baixa uma biblioteca (ex.: driver MySQL)   |
| `go fmt ./...`     | Formata o código no padrão da linguagem    |

## Links
- Tour of Go (oficial, interativo): https://go.dev/tour
- Go by Example (receitas curtas): https://gobyexample.com
- Effective Go (boas práticas): https://go.dev/doc/effective_go

---

# 📒 Anotações das Lições — Fase 1 (para estudar em casa)

> **🔖 VOCÊ PAROU NA LIÇÃO 4** (funções + métodos).
> **Próximos passos:**
> 1. Mini-desafio da Lição 4 → função `valorTotalEstoque(produtos []Produto) float64`
> 2. **Lição 5** — funções que retornam erro (`if err != nil`) → é o que destrava a Fase 2 (servidor HTTP)

## 🏠 Quando chegar em casa (ambiente no Mac)
- [ ] Instalar o Go: https://go.dev/dl — conferir com `go version`
- [ ] VS Code + extensão **"Go"** (da Google)
- [ ] Criar o repositório: `git init` (ou criar no GitHub e dar `git clone`)
- [ ] Na pasta do projeto: `go mod init estudosGo`
- [ ] Criar `main.go`, colar o "estado atual do código" (lá embaixo) e rodar `go run .`

## ⌨️ Comandos do dia a dia
| Comando | O que faz |
|---|---|
| `go mod init <nome>` | cria o `go.mod` (manifesto do projeto) |
| `go run .` | compila e roda na hora (pra testar) |
| `go build` | gera o `.exe` |
| `go fmt ./...` | formata o código no padrão Go |

---

## Lição 1 — Anatomia + tipagem estática
- `package main` + `func main()` (minúsculo e fixo) = ponto de entrada. É o `public static void main` do Java.
- `import "fmt"` traz a lib de impressão; usa-se `fmt.Println(...)`.
- **Maiúscula = público / minúscula = privado** — é o "public/private" do Go embutido na 1ª letra. Por isso é `fmt.Println` (P maiúsculo = exportado pela lib).
  - **Usando** algo de lib → escreve como o autor batizou (geralmente maiúsculo). **Definindo** o seu → maiúsculo só se quiser que outros pacotes vejam.
- `:=` declara **e** atribui (≈ `var` do Java); `=` só reatribui o que já existe.
- **Tipagem estática:** o tipo é fixo e o Go barra erro de tipo **na compilação** (ex.: pôr string num `int` não compila).

## Lição 2 — `struct` (o registro de produto)
- `struct` = os atributos de uma classe (só dados; **sem herança**).
- **O tipo vem DEPOIS do nome:** `Codigo string` (contrário do Java).
- Campos em **maiúsculo** porque vão virar **JSON** na Fase 2 (a lib só enxerga campo público).
- Sem `null`/`new`: os campos já nascem com **"valor zero"** (`0`, `""`, `false`).
- Criar: campo a campo, ou struct literal → `p := Produto{Codigo: "001", Nome: "Coca", Preco: 5.99, Quantidade: 4}`.

## Lição 3 — `slice` (a lista / o "browse")
- `[]Produto` = lista dinâmica = `ArrayList<>` (Java) / `aArraBrow` (Harbour).
- Índice começa em **0** ⚠️ (Harbour começa em 1).
- `len(lista)` = tamanho. `append(lista, item)` adiciona — **sempre reatribuindo**: `lista = append(lista, x)`.
- `for i, p := range lista` → `i` = índice, `p` = cópia do item. **Isso é o browse.**
- O **`_` (blank identifier):** o `range` SEMPRE devolve `(índice, valor)`. Se não usar o índice, descarta com `_` (Go **proíbe variável não usada**). Cuidado: `for p := range lista` te dá o ÍNDICE, não o item!

## Lição 4 — funções e métodos
- Função: `func nome(parametros) TipoRetorno { ... }` (o tipo de retorno vem **no fim**).
- Método (com *receiver*): `func (p Produto) ValorTotal() float64 { ... }` → chama com `p.ValorTotal()`.
- Go **não converte número automaticamente:** `p.Preco * float64(p.Quantidade)` (precisa do `float64(...)`).
- **Sem getter/setter automático:** acessa o campo direto (`p.Codigo`). Só cria método quando há lógica de verdade. Getter idiomático é `Preco()` (sem "Get"); setter usa ponteiro: `func (p *Produto) SetPreco(...)`.
- 🔑 `listarProdutos() []Produto` hoje devolve lista fixa; na **Fase 3** vai ler do MySQL com a **mesma assinatura**. Isso separa "de onde vêm os dados" de "quem usa" — é o esqueleto da camada de repositório (≈ `oBase:AbreTabe('PRODUTOS')` do Consys).

### Estado atual do `main.go`
```go
package main

import "fmt"

type Produto struct {
	Codigo     string
	Nome       string
	Preco      float64
	Quantidade int
}

// MÉTODO: pertence ao Produto (receiver "(p Produto)")
func (p Produto) ValorTotal() float64 {
	return p.Preco * float64(p.Quantidade)
}

// FUNÇÃO: "busca" os produtos. Hoje fixo; na Fase 3 virá do MySQL.
func listarProdutos() []Produto {
	return []Produto{
		{Codigo: "001", Nome: "Pepsi", Preco: 3.99, Quantidade: 11},
		{Codigo: "002", Nome: "Coca", Preco: 3.99, Quantidade: 52},
		{Codigo: "003", Nome: "Guarana", Preco: 3.99, Quantidade: 31},
	}
}

func main() {
	produtos := listarProdutos()

	for i, p := range produtos {
		fmt.Println(i+1, "-", p.Nome, "R$", p.Preco, "| estoque: R$", p.ValorTotal())
	}
}
```

### ⏳ Mini-desafio pendente (Lição 4)
Criar `func valorTotalEstoque(produtos []Produto) float64` que percorre a lista, soma o `ValorTotal()` de cada produto e devolve o total geral. Chamar na `main` e imprimir.
Dica: comece com `total := 0.0` e vá somando dentro do `for`.
