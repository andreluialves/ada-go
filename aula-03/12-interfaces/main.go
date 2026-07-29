package main

import "fmt"

type Saudador interface {
	Saudar() string
}

type Paciente struct {
	Nome string
}

func (p *Paciente) Saudar() string {
	if p == nil {
		return "paciente nil, mas com tipo concreto definido"
	}

	return "Ola, " + p.Nome
}

func main() {
	var paciente *Paciente
	fmt.Println("ponteiro paciente == nil?", paciente == nil)

	var saudador Saudador
	fmt.Println("interface saudador == nil?", saudador == nil)

	saudador = paciente
	fmt.Println("interface depois de receber paciente == nil?", saudador == nil)
	fmt.Println(saudador.Saudar())
}
