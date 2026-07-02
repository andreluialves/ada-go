// Exemplo de como usar um channel simples
package main

import "fmt"

func exemploChannelSimples() {
	ch := make(chan string)

	go func() {
		ch <- "Olá, eu vim de uma goroutine"
	}()

	mensagem := <-ch
	fmt.Println(mensagem)
}

func main() {
	ch := make(chan string, 5)
	fmt.Println(ch)
	exemploChannelSimples()
}
