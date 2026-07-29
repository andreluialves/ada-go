package shipping

func hasFreeShipping(totalCents int) bool {
	return totalCents >= 20_000
}
