package order

import (
	"errors"
	"testing"
)

func TestTransitionOrder(t *testing.T) {
	testCases := []struct {
		name      string
		current   status
		requested action
		want      status
		wantErr   error
	}{
		{name: "paga pedido aberto", current: statusOpen, requested: actionPay, want: statusPaid},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := transitionOrder(tc.current, tc.requested)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("erro = %v; esperado = %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("estado = %q; esperado = %q", got, tc.want)
			}
		})
	}
}
