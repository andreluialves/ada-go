package username

import "testing"

func TestNormalizeUsername(t *testing.T) {
	testCases := []struct {
		name  string
		value string
		want  string
	}{
		{name: "remove espaços", value: "  Luan  ", want: "luan"},
		{name: "converte para minúsculas", value: "ALUNA", want: "aluna"},
		{name: "mantém valor normalizado", value: "golang", want: "golang"},
	}

	for _, tc := range testCases {
		got := normalizeUsername(tc.value)
		if got != tc.want {
			t.Errorf("%s: username = %q; esperado = %q", tc.name, got, tc.want)
		}
	}
}
