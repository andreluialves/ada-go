// Embedding de interfaces
package main

import "fmt"

type Leitor interface {
	Ler() string
}

type Fechador interface {
	Fechar()
}

type LeitorFechador interface {
	Leitor
	Fechador
}

type ArquivoEmMemoria struct {
	Nome     string
	Conteudo string
	Fechado  bool
}

func (a *ArquivoEmMemoria) Ler() string {
	if a.Fechado {
		return "arquivo fechado"
	}
	return a.Conteudo
}

func (a *ArquivoEmMemoria) Fechar() {
	a.Fechado = true
}

func usarArquivo(arquivo LeitorFechador) {
	fmt.Println("conteudo:", arquivo.Ler())
	arquivo.Fechar()
	fmt.Println("depois de fechar:", arquivo.Ler())
}

func main() {
	arquivo := &ArquivoEmMemoria{
		Nome:     "pacientes.txt",
		Conteudo: "Luan, Ana, Joao",
	}

	usarArquivo(arquivo)
}
