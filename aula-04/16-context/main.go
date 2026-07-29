// WithValue
package main

import (
	"context"
	"fmt"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func logar(ctx context.Context, mensagem string) {
	requestID := ctx.Value(requestIDKey)

	fmt.Println("request_id:", requestID)
	fmt.Println("mensagem:", mensagem)
}

func processarPedido(ctx context.Context) {
	logar(ctx, "Processando pedido")
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, requestIDKey, "abc-123")

	processarPedido(ctx)
}
