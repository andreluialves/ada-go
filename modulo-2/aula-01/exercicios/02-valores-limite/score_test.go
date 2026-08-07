package score

import "testing"

func TestIsValidScore(t *testing.T) {
	if got := isValidScore(50); !got {
		t.Errorf("isValidScore(50) = %v; esperado = true", got)
	}
}
