package withdraw

import "errors"

var (
	errInvalidAmount       = errors.New("invalid withdrawal amount")
	errInsufficientBalance = errors.New("insufficient balance")
)

func withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return balance, errInvalidAmount
	}

	if amount > balance {
		return balance, errInsufficientBalance
	}

	return balance - amount, nil
}
