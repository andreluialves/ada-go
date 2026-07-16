package ports

import "testing"

func TestIsValidPort(t *testing.T) {
	testCases := []struct {
		name string
		port int
		want bool
	}{
		{
			name: "zero é inválida",
			port: 0,
			want: false,
		},
		{
			name: "limite inferior é válido",
			port: 1,
			want: true,
		},
		{
			name: "porta comum é válida",
			port: 8080,
			want: true,
		},
		{
			name: "limite superior é válido",
			port: 65535,
			want: true,
		},
		{
			name: "acima do limite é inválida",
			port: 65536,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidPort(tc.port)

			if got != tc.want {
				t.Errorf(
					"isValidPort(%d) = %v, want %v",
					tc.port,
					got,
					tc.want,
				)
			}
		})
	}
}
