package subtotal

import "testing"

func TestCalcularSubtotal(t *testing.T) {
	// arrange - preparação
	unitPriceCents := 500
	quantity := 3
	want := 1500

	// Act - Executar
	got := calcularSubtotal(unitPriceCents, quantity)

	// Assertion
	if got != want {
		t.Fatalf("calcularSubtotal(%d,%d) = %d, want %d",
			unitPriceCents,
			quantity,
			got,
			want,
		)
	}
}
