package calculator

import (
	"testing"
)

func TestService(t *testing.T) {
	result := Calculate("10 + 5")
	if got, expected := result, "15.00"; got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
