# Exercício 6 — testando resultados e erros

## Objetivo

Testar funções que retornam um resultado e um erro, usando `errors.Is` para comparar erros sentinela.

Regras de `withdraw`:

- o valor do saque deve ser maior que zero;
- o saque não pode ultrapassar o saldo;
- quando ocorre erro, o saldo deve permanecer inalterado;
- sacar exatamente o saldo disponível é permitido.

## Tarefas

Adicione à tabela os seguintes casos:

| Cenário | Saldo | Saque | Saldo esperado | Erro esperado |
|---|---:|---:|---:|---|
| valor zero | `100` | `0` | `100` | `errInvalidAmount` |
| valor negativo | `100` | `-10` | `100` | `errInvalidAmount` |
| saldo insuficiente | `100` | `150` | `100` | `errInsufficientBalance` |
| saque de todo o saldo | `100` | `100` | `0` | `nil` |

Execute `go test -v` e observe que cada cenário possui seu próprio subtest.

