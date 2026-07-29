// WithCancel
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker recebeu cancelamento:", ctx.Err())
			return
		default:
			fmt.Println("Worker trabalhando...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)

	fmt.Println("Main cancelando o contexto...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main finalizada")
}
