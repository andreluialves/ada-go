# Exercício 5 — subtests

## Objetivo

Transformar cada linha de uma tabela em um cenário identificado e executável separadamente.

O teste existente já é orientado por tabela, mas todos os casos ainda pertencem diretamente a `TestNormalizeUsername`.

## Tarefas

1. Envolva a execução de cada caso com:

```go
t.Run(tc.name, func(t *testing.T) {
	// Act e Assert
})
```

2. Execute `go test -v` e observe a hierarquia dos subtests.
3. Execute somente o cenário que remove espaços:

```sh
go test -v -run 'TestNormalizeUsername/remove_espaços'
```

4. Como o nome já aparecerá no caminho do subtest, experimente remover `tc.name` da mensagem de erro.

