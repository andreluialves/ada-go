package main

import "fmt"

type Registrador interface {
	Registrar(mensagem string)
}

type FuncaoRegistradora func(mensagem string)

func (f FuncaoRegistradora) Registrar(mensagem string) {
	f(mensagem)
}

func registrarNoConsole(mensagem string) {
	fmt.Println("console:", mensagem)
}

func executarConsulta(registrador Registrador, cpf string) {
	registrador.Registrar("consultando paciente com CPF " + cpf)
}

func main() {
	registrador := FuncaoRegistradora(registrarNoConsole)

	executarConsulta(registrador, "111")
}
