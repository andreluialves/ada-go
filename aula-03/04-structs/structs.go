package main

import (
	"encoding/json"
	"fmt"
)

type Paciente struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
	CPF   string `json:"cpf"`
}

func main() {
	paciente := Paciente{
		Nome:  "Luan",
		Idade: 34,
		CPF:   "0123456789",
	}
	fmt.Println("paciente sem tags:", paciente)
	jsonPaciente, err := json.Marshal(paciente)
	if err != nil {
		fmt.Println("erro ao converter para JSON:", err)
	}
	fmt.Println("paciente com tags:", string(jsonPaciente))
}
