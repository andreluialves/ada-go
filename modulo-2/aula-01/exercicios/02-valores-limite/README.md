# Exercício 2 — valores de fronteira

## Objetivo

Mostrar que um teste comum e 100% de cobertura de instruções ainda podem esconder um bug.

O contrato de `isValidScore` diz que notas entre 0 e 100, inclusive, são válidas.

## Tarefas

1. Execute `go test -cover` e observe o resultado.
2. Adicione testes para os seguintes valores:

| Entrada | Esperado |
|---:|---:|
| `-1` | `false` |
| `0` | `true` |
| `50` | `true` |
| `100` | `true` |
| `101` | `false` |

3. Execute os testes novamente e identifique as fronteiras defeituosas.
4. Corrija `isValidScore` sem alterar o contrato.
5. Execute `go test -v -cover`.

Não corrija a função antes de observar o teste falhar.

