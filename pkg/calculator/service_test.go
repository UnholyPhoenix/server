package calculator

import (
	"errors"
	"testing"
)

func TestService(t *testing.T) {
	service := New()

	tests := []struct {
		name        string
		requestExpr string
		result      float64
		err         error
	}{
		{
			"Expression not set",
			"",
			0.0,
			errors.New("Error occured. Please write an valid expression"),
		},
		{
			"Expression for adding",
			"10 + 5",
			15.00,
			nil,
		},
		{
			"Expression for subtraction",
			"10 - 5",
			5.00,
			nil,
		},
		{
			"Expression for multiply",
			"10 * 5",
			50.00,
			nil,
		},
		{
			"Expression for dividing",
			"10 / 5",
			2.00,
			nil,
		},
		{
			"Expression with wrong method",
			"10 | 5",
			0.0,
			errors.New("Invalid operator. Allowed operators are '+', '-', '*' and '/'"),
		},
		{
			"Expression for dividing with zero",
			"10 / 0",
			0.0,
			errors.New("Not possible dividing with zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Evaluate(tt.requestExpr)
			if expect, got, gotError := tt.result, result, err; expect != got {
				t.Errorf("expected '%v', got '%v', error '%v'", expect, got, gotError)
			}
		})
	}
}
