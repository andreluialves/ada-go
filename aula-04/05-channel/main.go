// Exemplo de receber value, ok com channel
package main

import "fmt"

type Resultado struct {
	Status string
	Erro   error
}

func exemploValorEOkComStruct() {
	ch := make(chan Resultado)

	go func() {
		ch <- Resultado{
			Status: "ok",
			Erro:   nil,
		}

		close(ch)
	}()

	resultado, ok := <-ch
	fmt.Println("Resultado:", resultado)
	fmt.Println("Canal aberto?", ok)
	//ch <- Resultado{Status: "ok"}

	resultado, ok = <-ch
	fmt.Println("Resultado:", resultado)
	fmt.Println("Canal aberto?", ok)
}

func main() {
	exemploValorEOkComStruct()
}
