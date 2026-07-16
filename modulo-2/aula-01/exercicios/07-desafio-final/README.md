# Exercício 7 — desafio final: transições de pedido

## Objetivo

Combinar contrato, estados, table-driven tests, subtests, erros e cobertura.

Um pedido pode estar `aberto`, `pago` ou `cancelado`. As regras são:

| Estado atual | Ação | Resultado | Erro |
|---|---|---|---|
| aberto | pagar | pago | `nil` |
| aberto | cancelar | cancelado | `nil` |
| pago | cancelar | cancelado | `nil` |
| pago | pagar | pago | `errAlreadyPaid` |
| cancelado | pagar | cancelado | `errInvalidTransition` |
| cancelado | cancelar | cancelado | `errAlreadyCancelled` |
| qualquer estado | ação desconhecida | estado anterior | `errInvalidAction` |

## Tarefas

1. O teste inicial cobre somente o caminho mais comum. Adicione os demais cenários à tabela.
2. Execute `go test -v`.
3. Use o nome do subtest que falhar para localizar o comportamento incorreto.
4. Corrija a implementação sem alterar o contrato.
5. Execute:

```sh
go test -v -cover
```

6. Gere um relatório de cobertura:

```sh
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Use a cobertura para procurar caminhos ausentes, não como prova de que todos os comportamentos estão corretos.

