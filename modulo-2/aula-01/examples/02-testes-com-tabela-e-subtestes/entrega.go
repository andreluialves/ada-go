package entrega

func calcularTaxaEntrega(distanceKm int) int {
	switch {
	case distanceKm <= 0:
		return 0
	case distanceKm <= 5:
		return 500
	case distanceKm <= 15:
		return 1_000
	default:
		return 2_000
	}
}
