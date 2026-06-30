// Method Set: Diferença de Receiver value e Receiver Pointer
package main

import "fmt"

type Exibidor interface {
	Exibir() string
}

type Incrementador interface {
	Incrementar()
}

type Contador struct {
	Valor int
}

func (c Contador) Exibir() string {
	return fmt.Sprintf("valor atual: %d", c.Valor)
}

func (c *Contador) Incrementar() {
	c.Valor++
}

func mostrarValor(exibidor Exibidor) {
	fmt.Println(exibidor.Exibir())
}

func aumentarValor(incrementador Incrementador) {
	incrementador.Incrementar()
}

func main() {
	contador := Contador{Valor: 10}
	mostrarValor(contador)
	mostrarValor(&contador)
	aumentarValor(&contador)

	// Esta linha nao compila, porque Incrementar tem receiver ponteiro:
	aumentarValor(&contador)
}
