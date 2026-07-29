# Exercício 3 — testes repetidos

## Objetivo

Perceber que os dados mudam, mas a lógica do teste continua igual, e então refatorar para um table-driven test.

Pedidos com valor igual ou superior a 20.000 centavos possuem frete grátis.

## Tarefas

1. Execute `go test -v` e confirme que os cinco testes passam.
2. Observe quais trechos se repetem em todas as funções.
3. Substitua as cinco funções por uma única função `TestHasFreeShipping`.
4. Crie uma slice de structs chamada `testCases` com os campos `name`, `totalCents` e `want`.
5. Percorra a tabela usando `for _, tc := range testCases`.
6. Ainda não use `t.Run`; ele será praticado no exercício 5.
7. Inclua `tc.name` na mensagem de erro e execute `go test -v` novamente.

O comportamento final deve continuar exatamente igual ao dos cinco testes originais.

