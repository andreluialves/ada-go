package main

import (
	"fmt"

	"github.com/luanlouzada/ada-go/aula-03/03-metodos-pacote/pacote"
)

type Paciente struct {
	Nome  string
	Idade uint
	CPF   string
}

func (p *Paciente) atualizaNome(novoNome string) {
	p.Nome = novoNome
}

func atualizaNome(p *Paciente, novoNome string) {
	p.Nome = novoNome
}

func (e *pacote.Exemplo) atualizaValor(novoValor int) {
	e.Valor = novoValor
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
	exemplo := pacote.Exemplo{
		Valor: 10,
		Texto: "exemplo",
	}
	fmt.Println(exemplo)
}
