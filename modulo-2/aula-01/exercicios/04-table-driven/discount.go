package discount

func discountPercent(customerYears int) int {
	if customerYears < 2 {
		return 0
	}

	if customerYears < 5 {
		return 5
	}

	return 10
}
