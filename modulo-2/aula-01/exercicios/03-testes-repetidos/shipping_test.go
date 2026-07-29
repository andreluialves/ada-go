package shipping

import "testing"

func TestHasFreeShippingForZero(t *testing.T) {
	if hasFreeShipping(0) {
		t.Error("pedido de 0 centavos não deveria possuir frete grátis")
	}
}

func TestHasFreeShippingBelowLimit(t *testing.T) {
	if hasFreeShipping(19_999) {
		t.Error("pedido abaixo de 20.000 centavos não deveria possuir frete grátis")
	}
}

func TestHasFreeShippingAtLimit(t *testing.T) {
	if !hasFreeShipping(20_000) {
		t.Error("pedido de 20.000 centavos deveria possuir frete grátis")
	}
}

func TestHasFreeShippingAboveLimit(t *testing.T) {
	if !hasFreeShipping(20_001) {
		t.Error("pedido acima de 20.000 centavos deveria possuir frete grátis")
	}
}

func TestHasFreeShippingForLargeOrder(t *testing.T) {
	if !hasFreeShipping(50_000) {
		t.Error("pedido de 50.000 centavos deveria possuir frete grátis")
	}
}
