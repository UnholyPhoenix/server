package calculator

import (
	"testing"
)

func TestService(t *testing.T) {
	result := Calculate("10 + 5")
	if got, expected := result, "15.00"; got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	result = Calculate("10 - 5")
	if got, expected := result, "5.00"; got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	result = Calculate("10 * 5")
	if got, expected := result, "50.00"; got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}

	result = Calculate("10 / 5")
	if got, expected := result, "2.00"; got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
