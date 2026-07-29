package entrega

import "testing"

func TestCalcularTaxaEntrega(t *testing.T) {
	//arrange
	testCases := []struct {
		name       string
		distanceKm int
		want       int
	}{
		{
			name:       "distância negativa",
			distanceKm: -1,
			want:       0,
		},
		{
			name:       "distancia igual a 0",
			distanceKm: 0,
			want:       0,
		},
		{
			name:       "distancia igual maior que 0 ou menor igual a 5",
			distanceKm: 1,
			want:       500,
		},
		{
			name:       "distancia igual maior que 0 ou menor igual a 5",
			distanceKm: 5,
			want:       500,
		},
		{
			name:       "caso nenhum outro",
			distanceKm: 16,
			want:       2_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calcularTaxaEntrega(tc.distanceKm)

			if got != tc.want {
				t.Errorf("calcularTaxaEntrega(%d) = %d e want %d",
					tc.distanceKm,
					got,
					tc.want,
				)
			}
		})
	}
}
