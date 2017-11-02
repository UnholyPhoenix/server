package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// Service greets name
type Service interface {
	Evaluate(expression string) (float64, error)
}

// Calculate struct for calculating
type Calculate struct{}

// Evaluate expression and returns the value
func (c Calculate) Evaluate(expression string) (float64, error) {
	// Splits input by space
	statment := strings.Fields(expression)

	if len(statment) != 3 {
		return 0, errors.New("Error occured. Please write an valid expression")
	}

	// Expression parts
	firstNumber := statment[0]
	operator := statment[1]
	secondNumber := statment[2]

	firstN, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		return 0, errors.New("Error occured. Please write an valid expression")
	}

	secondN, err := strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		return 0, errors.New("Error occured. Please write an valid expression")
	}
	var result float64
	switch {
	case operator == "+":
		result = firstN + secondN

	case operator == "-":
		result = firstN - secondN
	case operator == "*":
		result = firstN * secondN
	case operator == "/" && secondN == 0.00:
		return 0.0, errors.New("Not possible dividing with zero")
	case operator == "/":
		result = firstN / secondN
	default:
		return 0, errors.New("Invalid operator. Allowed operators are '+', '-', '*' and '/'")
	}
	return result, nil
}

// New service
func New() Calculate {
	return Calculate{}
}
