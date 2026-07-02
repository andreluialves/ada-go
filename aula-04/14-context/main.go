// WithDeadline
package main

import (
	"context"
	"fmt"
	"time"
)

func operacaoComDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operação cancelada:", ctx.Err())
			return

		default:
			fmt.Println("Operação ainda dentro do prazo...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	fmt.Println("Deadline definida para:", deadline.Format("15:04:05"))

	//"15:04:05" // hora 24h
	// "03:04:05" // hora 12h com zero
	// "3:04:05 PM" // hora 12h com AM/PM
	// https://pkg.go.dev/time
	fmt.Println(deadline)
	operacaoComDeadline(ctx)
	fmt.Println("Main finalizada")
}
