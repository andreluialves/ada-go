// Exemplo de deadlock recebendo valor de um channel sem enviar valor
package main

import "fmt"

func main() {
	ch := make(chan string)

	// go func() {
	// 	ch <- "Olá, eu vim de uma goroutine"
	// }()
	// Sem enviar valor, vai dar deadlock
	msg := <-ch
	fmt.Println(msg)
}
