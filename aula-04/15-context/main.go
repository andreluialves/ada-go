// WithTimeout
package main

import (
	"context"
	"fmt"
	"time"
)

func chamadaExterna(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Chamada externa finalizada com sucesso")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("Iniciando chamada externa...")

	err := chamadaExterna(ctx)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Main finalizada com sucesso")
}
