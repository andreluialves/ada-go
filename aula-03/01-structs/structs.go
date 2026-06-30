package main

import (
	"encoding/json"
	"fmt"
)

type PacienteNome string

type PacienteIdade uint

type PacienteCpf string

type Paciente struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
	CPF   string `json:"cpf" xml:"cpf"`
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
	jsonPaciente, err := json.Marshal(paciente)
	if err != nil {
		fmt.Println("erro ao converter para JSON:", err)
	}
	fmt.Println("paciente com tags:", string(jsonPaciente))
}
