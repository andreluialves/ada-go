// Caso de uso real consulta 2 api de terceiros e retorna a primeira resposta
package main

import (
	"context"
	"fmt"
	"time"
)

type Cotacao struct {
	Provedor string
	Valor    float64
}

func consultarProvedor(ctx context.Context, nome string, demora time.Duration, valor float64, ch chan<- Cotacao) {
	select {
	case <-time.After(demora):
		select {
		case ch <- Cotacao{
			Provedor: nome,
			Valor:    valor,
		}:
		case <-ctx.Done():
			return
		}

	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan Cotacao, 2)

	go consultarProvedor(ctx, "Provedor A", 1500*time.Millisecond, 5.10, ch)
	go consultarProvedor(ctx, "Provedor B", 800*time.Millisecond, 5.08, ch)

	select {
	case cotacao := <-ch:
		fmt.Println("Usei cotação do:", cotacao.Provedor)
		fmt.Println("Valor:", cotacao.Valor)
		return

	case <-ctx.Done():
		fmt.Println("Nenhum provedor respondeu a tempo")
		return
	}
}
