# 🐹 Estudos de Golang — De Harbour para a Web

Este repositório documenta minha jornada de aprendizado na linguagem **Go**, migrando de uma base em **Harbour/xHarbour** para a construção de APIs modernas e escaláveis.

## 🎯 Objetivo

Transformar o conhecimento em sistemas desktop (ERP, CRUD, SQL) em habilidades de desenvolvimento Backend, conectando o Go ao MySQL e criando interfaces web (HTML/JS).

---

## 🚀 Progresso do Roadmap

| Fase  | Descrição                           | Status          |
| :---- | :---------------------------------- | :-------------- |
| **0** | Configuração do Ambiente            | ✅ Concluído    |
| **1** | Fundamentos da Linguagem (Terminal) | 🟡 Em andamento |
| **2** | Primeiro Servidor HTTP              | 📅 Planejado    |
| **3** | Conexão MySQL (check_fl02)          | 📅 Planejado    |
| **4** | Frontend (AJAX / Fetch)             | 📅 Planejado    |
| **5** | CRUD Completo                       | 📅 Planejado    |

> Veja o plano detalhado em: [ROADMAP.md](./ROADMAP.md)

---

## 🛠️ Tecnologias e Ferramentas

- **Linguagem:** Go (Golang) v1.26+
- **Editor:** VS Code + Go Extension
- **Banco de Dados:** MySQL (check_fl02)
- **Protocolo:** REST / JSON

---

## 📖 Lições Aprendidas (Resumo)

### 1. Tipagem e Sintaxe

- O Go utiliza **tipagem estática** (erros pegos em tempo de compilação).
- Operador `:=` para declaração rápida e `=` para atribuição.
- Visibilidade definida pela capitalização: `Maiúscula` (Público/Exportado) vs `minúscula` (Privado).

### 2. Estruturas de Dados

- **Structs:** Agrupamento de dados (equivalente a registros/classes sem herança).
- **Slices:** Arrays dinâmicos. Importante: o índice começa em **0**.
- **Maps:** Coleções chave-valor.

### 3. Tratamento de Erros

- Diferente do Harbour (`BEGIN SEQUENCE`), o Go utiliza o retorno explícito de erros: `if err != nil`.

---

## 💻 Como Rodar

Para executar os exemplos atuais:

```bash
go run main.go
```

Para gerar o executável:

```bash
go build -o estudosGo
```

---

## 🔗 Links Úteis

- [Tour of Go (Oficial)](https://go.dev/tour)
- [Go by Example](https://gobyexample.com)
- [Documentação Oficial](https://go.dev/doc/)

---

_Desenvolvido com foco na transição tecnológica Consys/Harbour -> Go._
