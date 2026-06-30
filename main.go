package main

import "fmt"

type Paciente struct {
	Nome  string
	Idade uint
	CPF   string
}

type PacienteEstrangeiro struct {
	DadosPaciente Paciente
	Passaporte    string
}

func main() {
	pacienteEstrangeiro := PacienteEstrangeiro{
		DadosPaciente: Paciente{
			Nome:  "João",
			Idade: 34,
			CPF:   "0123456789",
		},
		Passaporte: "1234567890",
	}
	fmt.Println(pacienteEstrangeiro.DadosPaciente.Nome)
	fmt.Println(pacienteEstrangeiro.DadosPaciente.Idade)
	fmt.Println(pacienteEstrangeiro.DadosPaciente.CPF)
	fmt.Println(pacienteEstrangeiro.Passaporte)
}
