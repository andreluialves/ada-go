## Exercicio 1 - Refatorando o processador de pedidos

Objetivo: transformar regras soltas dentro da `main` em funcoes pequenas e reutilizaveis.

Partindo deste codigo:

```go
package main

import "fmt"

func main() {
	total := 350.0
	items := 3

	if items <= 0 || total <= 0 {
		fmt.Println("pedido rejeitado")
		return
	}

	discount := total * 0.10
	finalTotal := total - discount

	fmt.Println("pedido aprovado")
	fmt.Println("desconto:", discount)
	fmt.Println("total final:", finalTotal)
}
```

Crie as funcoes:

```go
func validarPedido(items int, total float64) error
func calcularDesconto(total float64) float64
func processarPedido(items int, total float64) (float64, float64, error)
```

Regras:

- `validarPedido` deve retornar erro se `items <= 0`.
- `validarPedido` deve retornar erro se `total <= 0`.
- Se `items <= 0` e `total <= 0` ao mesmo tempo, retorne os dois erros usando `errors.Join`.
- Crie erros sentinela, por exemplo `ErrItemsInvalidos` e `ErrTotalInvalido`.
- `calcularDesconto` deve aplicar:
  - 15% se `total >= 500`
  - 10% se `total >= 200`
  - 0% caso contrario
- `processarPedido` deve retornar `desconto`, `totalFinal` e `error`.
- A `main` deve tratar o erro antes de imprimir o resultado.
- Na `main`, use `errors.Is` para identificar qual regra falhou.

Casos para testar manualmente:

```txt
items=3 total=350   -> desconto 35, total final 315
items=2 total=600   -> desconto 90, total final 510
items=1 total=100   -> desconto 0, total final 100
items=0 total=300   -> erro de itens
items=2 total=-10   -> erro de total
items=0 total=-10   -> erro de itens e erro de total
```

Pergunta para responder:

- Por que `processarPedido` retorna tres valores em vez de imprimir tudo dentro dela?