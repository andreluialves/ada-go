package main

import "fmt"

type Pagamento struct {
	ID     string
	Valor  int
	Status bool
}

func NovoPagamento(id string, valor int, status bool) *Pagamento {
	return &Pagamento{
		ID:     id,
		Valor:  valor,
		Status: status,
	}
}

func (p *Pagamento) AtualizarStatus() {
	p.Status = true
}
func AtualizarStatus(p *Pagamento) {
	p.Status = true
}

func main() {
	pagamento := Pagamento{
		ID:     "123",
		Valor:  100,
		Status: false,
	}
	pagamento2 := pagamento
	fmt.Printf("Pagamento: %p\n", &pagamento)
	fmt.Printf("Pagamento2: %p\n", &pagamento2)
	pagamento3 := NovoPagamento("456", 200, false)
	fmt.Printf("Pagamento3: %p\n", pagamento3)
	fmt.Println("Antes Pagamento3:", pagamento3)
	pagamento3.AtualizarStatus()
	fmt.Println("Pagamento3:", pagamento3)
}
