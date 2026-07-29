package withdraw

import (
	"errors"
	"testing"
)

func TestWithdraw(t *testing.T) {
	testCases := []struct {
		name        string
		balance     int
		amount      int
		wantBalance int
		wantErr     error
	}{
		{name: "saque válido", balance: 100, amount: 40, wantBalance: 60},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotBalance, err := withdraw(tc.balance, tc.amount)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("erro = %v; esperado = %v", err, tc.wantErr)
			}
			if gotBalance != tc.wantBalance {
				t.Errorf("saldo = %d; esperado = %d", gotBalance, tc.wantBalance)
			}
		})
	}
}
