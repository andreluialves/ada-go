# Exercício: Sistema de Notas Escolares

## Objetivo

Crie um programa em Go que gerencia notas de um aluno pelo terminal.

O exercício deve ser feito **sem usar `struct` nem ponteiros**. Use apenas o que já vimos em aula: variáveis, `slice`, `for`, `if`, `switch`, funções e `fmt`.

## Menu

O programa deve ficar em loop até o usuário escolher sair:

```
--- Sistema de Notas ---
1 - Adicionar nota
2 - Ver relatório
0 - Sair
```

## Regras

### 1. Adicionar nota

- Peça uma nota entre **0 e 10**.
- Se a nota for inválida, mostre: `Nota inválida! Use um valor entre 0 e 10.`
- Se for válida, guarde no `slice` e mostre: `Nota adicionada!`

### 2. Ver relatório

- Se não houver notas, mostre: `Nenhuma nota cadastrada.`
- Caso contrário, exiba:
  - todas as notas cadastradas
  - a **média**
  - a **maior nota**
  - a **menor nota**
  - o **status**: `Aprovado` se a média for maior ou igual a 7, senão `Reprovado`

Exemplo:

```
--- Relatório ---
Nota 1: 8.0
Nota 2: 6.0
Nota 3: 10.0
Média: 8.0
Maior nota: 10.0
Menor nota: 6.0
Status: Aprovado
```

### 3. Sair

- Mostre `Até mais!` e encerre o programa.

### 4. Opção inválida

- Mostre: `Opção inválida`

## Funções sugeridas

Separe a lógica em funções para facilitar os testes:

| Função | O que faz |
|--------|-----------|
| `notaValida(nota float64) bool` | retorna `true` se a nota estiver entre 0 e 10 |
| `calcularMedia(notas []float64) float64` | calcula a média das notas |
| `encontrarMaior(notas []float64) float64` | retorna a maior nota |
| `encontrarMenor(notas []float64) float64` | retorna a menor nota |
| `estaAprovado(media float64) bool` | retorna `true` se a média for >= 7 |

## Testes

Crie o arquivo `main_test.go` e escreva testes para as funções acima.

Sugestões de casos:

- nota válida e inválida
- cálculo de média com várias notas
- média com lista vazia
- maior e menor nota
- aprovação e reprovação

Para rodar:

```bash
go test -v ./...
```

## Como executar

```bash
go run .
```

## Entrega

1. Implemente a solução em `main.go`
2. Crie os testes em `main_test.go`
3. Confirme que `go test -v ./...` passa
4. Faça commit e envie para o repositório
