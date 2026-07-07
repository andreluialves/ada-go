# Exemplo com sqlc

Este exemplo mostra o `sqlc` gerando codigo Go a partir de SQL escrito manualmente.

Instalar o sqlc:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1
```

Gerar codigo:

```bash
cd examples/sqlc
sqlc generate
cd ../..
```

Rodar exemplo:

```bash
go run ./examples/sqlc
```

Neste exemplo, o `main.go` cria e recria a tabela `users_sqlc_demo`.
Assim ele nao mexe na tabela `users` usada nos outros exemplos.

Em uma aplicacao real, a tabela viria de migration, nao de `DROP TABLE` dentro
do codigo Go.

Testar todos os exemplos:

```bash
go test ./...
```

Ideia principal:

```text
voce escreve SQL
sqlc gera structs e metodos Go tipados
o codigo Go chama os metodos gerados
```
