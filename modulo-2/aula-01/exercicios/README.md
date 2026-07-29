# Exercícios — testes em Go

Os exercícios seguem uma ordem progressiva. Faça um por vez: leia o contrato, execute os testes iniciais e só depois altere o código.

| Exercício | Tema principal |
|---|---|
| `01-primeiro-teste` | Estrutura de um teste e Arrange, Act, Assert |
| `02-valores-limite` | Fronteiras e a limitação da cobertura |
| `03-testes-repetidos` | Identificar repetição e refatorar para table-driven test |
| `04-table-driven` | Completar uma tabela de cenários |
| `05-subtests` | Transformar cada caso em um subtest com `t.Run` |
| `06-testando-erros` | Resultados, erros sentinela e `errors.Is` |
| `07-desafio-final` | Estados, table-driven tests, subtests e cobertura |

## Comandos úteis

Execute os comandos dentro da pasta de cada exercício:

```sh
go test
go test -v
go test -cover
```

Para executar todos os exercícios a partir da pasta `aula-01`:

```sh
go test ./...
```

Os testes iniciais passam. Em alguns exercícios, adicionar os casos pedidos revelará bugs intencionais. Nesses casos, primeiro veja o teste falhar e somente depois corrija a implementação.

