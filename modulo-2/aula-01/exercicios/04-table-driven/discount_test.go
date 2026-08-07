package discount

import "testing"

func TestDiscountPercent(t *testing.T) {
	testCases := []struct {
		name          string
		customerYears int
		want          int
	}{
		{name: "cliente novo", customerYears: 0, want: 0},
		{name: "dois anos", customerYears: 2, want: 5},
		{name: "cinco anos", customerYears: 5, want: 10},
	}

	for _, tc := range testCases {
		got := discountPercent(tc.customerYears)
		if got != tc.want {
			t.Errorf("%s: desconto = %d; esperado = %d", tc.name, got, tc.want)
		}
	}
}
