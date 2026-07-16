# Exercício 4 — completando um table-driven test

## Objetivo

Representar partições e fronteiras como linhas de uma tabela de testes.

O desconto depende do tempo de relacionamento do cliente:

| Anos | Desconto |
|---:|---:|
| Menos de 2 anos | 0% |
| De 2 a 4 anos | 5% |
| 5 anos ou mais | 10% |

Valores negativos também devem produzir 0%.

## Tarefas

Complete `testCases` com estes valores:

| Entrada | Esperado |
|---:|---:|
| `-1` | `0` |
| `0` | `0` |
| `1` | `0` |
| `2` | `5` |
| `4` | `5` |
| `5` | `10` |
| `10` | `10` |

Depois execute:

```sh
go test -v
```

Observe que adicionar um cenário exige apenas uma nova linha, sem duplicar a lógica do teste.

