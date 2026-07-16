package desconto

import "errors"

var (
	errInvalidTotal    = errors.New("total must be greater than zero")
	errInvalidDiscount = errors.New("discount must be between 0 and 100")
)

func aplicarDesconto(totalCents, discountPercent int) (int, error) {
	if totalCents <= 0 {
		return totalCents, errInvalidTotal
	}

	if discountPercent < 0 || discountPercent > 100 {
		return totalCents, errInvalidDiscount
	}

	finalTotal := totalCents * (100 - discountPercent) / 100
	return finalTotal, nil
}
