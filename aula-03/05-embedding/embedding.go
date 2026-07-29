package main

import "fmt"

type Paciente struct {
	Nome  string
	Idade uint
	CPF   string
}

type PacienteEstrangeiro struct {
	Paciente
	Passaporte string
}

func (p *Paciente) atualizaNome(novoNome string) {
	p.Nome = novoNome
}

func atualizaNome(p *Paciente, novoNome string) {
	p.Nome = novoNome
}

func main() {
	paciente := Paciente{
		Nome:  "Luan",
		Idade: 34,
		CPF:   "0123456789",
	}
	paciente.atualizaNome("João")
	atualizaNome(&paciente, "Ana")
	fmt.Println(paciente)
	pacienteEstrangeiro := PacienteEstrangeiro{
		Paciente: Paciente{
			Nome:  "João",
			Idade: 34,
			CPF:   "0123456789",
		},
		Passaporte: "1234567890",
	}
	fmt.Println(pacienteEstrangeiro)
	pacienteEstrangeiro.atualizaNome("Maria")
	fmt.Println(pacienteEstrangeiro)
}
