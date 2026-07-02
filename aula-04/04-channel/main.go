// Exemplo de como fechar um channel
package main

import "fmt"

func exemploCloseChannel() {
	ch := make(chan string)

	go func() {
		ch <- "Mensagem 1"
		ch <- "Mensagem 2"
		ch <- "Mensagem 3"

		close(ch) // Se eu não fechar o channel o range vai esperar indefinidamente
	}()

	for mensagem := range ch {
		fmt.Println(mensagem)
	}
}

func main() {
	exemploCloseChannel()
}
