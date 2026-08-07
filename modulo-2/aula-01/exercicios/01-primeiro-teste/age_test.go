package age

import "testing"

func TestIsAdult(t *testing.T) {
	// Arrange
	age := 20

	// Act
	got := isAdult(age)

	// Assert
	if !got {
		t.Errorf("isAdult(%d) = %v; esperado = true", age, got)
	}
}
