// Exemplo de deadlock enviando valor para um channel com buffer cheio
package main

import "fmt"

// Tenho que ter algo consumindo o channel para enviar mais valores
func exemploBufferComConsumidor() {
	ch := make(chan string, 2)

	go func() {
		for msg := range ch {
			fmt.Println("Recebido:", msg)
		}
	}()

	ch <- "A"
	ch <- "B"
	ch <- "C"

	close(ch)
}

func main() {
	ch := make(chan string, 2)

	ch <- "A"
	ch <- "B"

	ch <- "C"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
