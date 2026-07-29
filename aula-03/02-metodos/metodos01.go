package main

import "fmt"

type Paciente struct {
	Nome  string
	Idade uint
	CPF   string
}

func (Paciente) fazNada() {
	fmt.Println("faz nada")
}

func main() {
	paciente := Paciente{
		Nome:  "Luan",
		Idade: 34,
		CPF:   "0123456789",
	}

	paciente.fazNada()
}
