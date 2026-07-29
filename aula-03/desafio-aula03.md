# Desafio Go: Dominio de um Servico de Pedidos

## Contexto

Voce foi chamado para construir o nucleo de um servico de pedidos para uma loja.

Nesta etapa nao existe HTTP, banco de dados, framework, JSON, Docker ou integracao externa.

O foco e a parte interna da aplicacao:

- dominio
- regras de negocio
- repositories em memoria
- interfaces
- services
- tratamento de erros

O `main.go` deve apenas montar alguns dados, chamar os services e imprimir resultados no terminal.

## Regras do desafio

- Nao implemente HTTP.
- Nao use banco de dados.
- Nao use framework.
- Use apenas a biblioteca padrao do Go.
- Nao use `panic` para regra de negocio.
- Quando algo der errado, retorne `error`.
- O service deve conversar com repositories por interfaces.
- O repository em memoria deve usar `map`.

## Problema

A loja possui produtos em estoque.

Um cliente pode criar um pedido com um ou mais produtos.

Quando um pedido e criado:

- o cliente precisa ter nome;
- o pedido precisa ter pelo menos um item;
- cada item precisa informar produto e quantidade;
- quantidade precisa ser maior que zero;
- produto precisa existir;
- produto precisa ter estoque suficiente;
- o estoque deve ser reduzido;
- o preco usado no pedido deve ser o preco do produto no momento da criacao;
- o pedido deve nascer como `PENDING`.

Depois de criado, um pedido pode ser:

- pago;
- cancelado.

Um pedido pago deve ficar como `PAID`.

Um pedido cancelado deve ficar como `CANCELED` e devolver os produtos ao estoque.

Pedidos pagos ou cancelados nao podem mudar de status novamente.

## Entidades esperadas

Crie os tipos que voce achar necessarios para representar este dominio.

No minimo, seu sistema deve conseguir representar:

- produto;
- item de pedido;
- pedido;
- status do pedido.

Um produto precisa ter informacoes suficientes para identificar o produto, exibir seu nome, saber seu preco e controlar o estoque.

Um item de pedido precisa guardar qual produto foi comprado, qual preco foi usado e qual quantidade foi pedida.

Um pedido precisa guardar identificacao, cliente, itens e status.

Voce tambem deve criar comportamentos no dominio. Exemplos:

- calcular total do pedido;
- reduzir estoque de um produto;
- devolver estoque de um produto;
- pagar pedido;
- cancelar pedido.

Use pointer receiver quando o metodo precisar alterar o valor.

## Erros esperados

Crie erros conhecidos para as principais falhas do dominio.

Seu codigo deve conseguir diferenciar, no minimo:

- produto nao encontrado;
- pedido nao encontrado;
- quantidade invalida;
- estoque insuficiente;
- cliente invalido;
- pedido vazio;
- mudanca de status invalida.

Use `errors.Is` quando fizer sentido comparar erros conhecidos.

Use `fmt.Errorf` com `%w` quando quiser adicionar contexto sem perder o erro original.

## Repositories

Crie repositories para produtos e pedidos.

Cada repository deve ter uma interface e uma implementacao em memoria.

O repository de produtos deve permitir:

- salvar produto;
- buscar produto por identificador;
- listar produtos.

O repository de pedidos deve permitir:

- salvar pedido;
- buscar pedido por identificador;
- listar pedidos.

A implementacao em memoria deve usar `map`.

Quando algo nao for encontrado, retorne o erro conhecido correspondente.

## Service de pedidos

Crie um service responsavel pelos casos de uso de pedido.

Ele deve receber os repositories por interface.

O service deve permitir:

- criar pedido;
- pagar pedido;
- cancelar pedido;
- buscar pedido por identificador;
- listar pedidos.

Na criacao do pedido, o service deve coordenar o fluxo:

- validar dados recebidos;
- buscar produtos;
- validar estoque;
- montar os itens;
- reduzir estoque;
- criar pedido;
- salvar pedido.

Uma dica importante: valide tudo que puder antes de alterar estoque. Assim voce evita deixar o sistema em um estado estranho quando um item do pedido falha no meio do processo.

## Geracao de IDs

Para simplificar, os IDs podem ser gerados em memoria.

Nao precisa usar UUID nem pacote externo.

Uma estrategia simples e manter um contador no service ou em algum gerador simples.

Exemplo de formato aceito:

```text
PED-001
PED-002
PED-003
```

## Estrutura do projeto

Voce pode organizar em varios arquivos dentro de `package main`.

Exemplo de separacao simples:

```text
pedidos/
  go.mod
  main.go
  product.go
  order.go
  errors.go
  repositories.go
  services.go
```

Se quiser separar em packages, tambem pode. Mas nao e obrigatorio.

O importante e nao deixar toda a regra de negocio dentro do `main.go`.

## Cenario minimo no main.go

No `main.go`, monte um fluxo que demonstre o sistema funcionando.

Cadastre alguns produtos:

```text
P001 - Notebook - 3500.00 - estoque 5
P002 - Mouse    - 80.00   - estoque 10
P003 - Teclado  - 180.00  - estoque 8
```

Depois demonstre:

1. Criacao de um pedido valido para `Ana` com `1 Notebook` e `2 Mouses`.
2. Impressao do pedido criado, status e total.
3. Impressao do estoque atualizado.
4. Pagamento do pedido.
5. Busca do pedido pago.
6. Tentativa de cancelar um pedido ja pago.
7. Tentativa de criar um pedido com estoque insuficiente.
8. Tentativa de criar um pedido invalido.

## Casos que precisam funcionar

- Criar pedido valido.
- Criar pedido sem cliente deve retornar erro.
- Criar pedido sem itens deve retornar erro.
- Criar pedido com produto inexistente deve retornar erro.
- Criar pedido com quantidade zero ou negativa deve retornar erro.
- Criar pedido com estoque insuficiente deve retornar erro.
- Ao criar pedido valido, estoque deve diminuir.
- Ao cancelar pedido pendente, estoque deve voltar.
- Pedido pago nao pode ser cancelado.
- Pedido cancelado nao pode ser pago.
- Buscar pedido inexistente deve retornar erro.

## Criterios de aceite

Seu projeto sera considerado completo se:

- compilar sem erros;
- rodar pelo terminal;
- tiver dominio, repositories, interfaces e service;
- tiver repository em memoria usando `map`;
- demonstrar pelo menos um fluxo feliz;
- demonstrar pelo menos tres fluxos de erro;
- tratar erros retornados, sem ignorar com `_`;
- nao tiver HTTP;
- nao tiver banco de dados.

## Roteiro sugerido

1. Crie o projeto com `go mod init pedidos`.
2. Modele as entidades do dominio.
3. Modele os erros conhecidos.
4. Defina as interfaces dos repositories.
5. Implemente os repositories em memoria.
6. Implemente o service de pedidos.
7. Monte o fluxo no `main.go`.
8. Rode o projeto pelo terminal.
9. Ajuste os casos de erro.

## Dicas

- Comece pelo dominio antes do service.
- O dominio deve concentrar regras que pertencem ao produto ou ao pedido.
- O service deve coordenar o caso de uso.
- O repository deve apenas guardar e buscar dados.
- Evite nomes genericos como `x`, `a`, `coisa` e `dados`.
- Prefira funcoes pequenas.
- Retorne cedo quando encontrar erro.
- Nao ignore erro retornado.
- Se um metodo altera a struct, pense em pointer receiver.
- Se um metodo apenas calcula algo, pense em receiver normal.

## Desafios bonus

Se terminar a parte obrigatoria, escolha um ou mais bonus.

### Bonus 1: Filtro com funcao como parametro

Crie uma forma de filtrar pedidos usando uma funcao recebida como parametro.

Use isso para listar:

- pedidos pagos;
- pedidos pendentes;
- pedidos acima de um determinado valor.

### Bonus 2: Desconto simples

Se o pedido passar de `5000.00`, aplique 10% de desconto.

Voce decide onde essa regra deve ficar.

### Bonus 3: Validacao agrupada

Ao validar um produto ou pedido, junte mais de um erro usando `errors.Join`.

Exemplos de problemas que poderiam aparecer juntos:

- produto sem nome;
- produto com preco invalido;
- produto com estoque negativo.

## Perguntas finais

Ao terminar, responda:

```text
1. Qual regra ficou no dominio?
2. Qual regra ficou no service?
3. Por que o service depende de interfaces?
4. O que mudaria se o repository deixasse de ser memoria e virasse banco de dados?
```
