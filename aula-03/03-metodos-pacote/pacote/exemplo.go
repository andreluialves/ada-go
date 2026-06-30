package pacote

type Exemplo struct {
	Valor int
	Texto string
}

func (e *Exemplo) atualizaValor(novoValor int) {
	e.Valor = novoValor
}
