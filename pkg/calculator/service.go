package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// Service greets name
type Service interface {
	Evaluate(expression string, kv *KV) (float64, error)
}

// Calculate struct for calculating
type Calculate struct{}

// Evaluate expression and returns the value
func (c Calculate) Evaluate(expression string, kv *KV) (float64, error) {
	// Splits input by space
	statment := strings.Fields(expression)

	if len(statment) != 3 {
		return 0, errors.New("Error occured. Please write an valid expression")
	}

	// Expression parts
	firstNumber := statment[0]
	operator := statment[1]
	secondNumber := statment[2]

	var firstN, secondN float64
	var err error

	if strings.ContainsAny("$", firstNumber) {
		firstN, err = getNumberFromMemory(firstNumber, kv)
		if err != nil {
			return 0, errors.New("Error occured. Please write an valid expression")
		}
	} else {
		firstN, err = strconv.ParseFloat(firstNumber, 64)
		if err != nil {
			return 0, errors.New("Error occured. Please write an valid expression")
		}
	}

	if strings.ContainsAny("$", secondNumber) {
		secondN, err = getNumberFromMemory(secondNumber, kv)
		if err != nil {
			return 0, errors.New("Error occured. Please write an valid expression")
		}
	} else {
		secondN, err = strconv.ParseFloat(secondNumber, 64)
		if err != nil {
			return 0, errors.New("Error occured. Please write an valid expression")
		}
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

func getNumberFromMemory(number string, kv *KV) (float64, error) {
	key, err := strconv.Atoi(strings.Split(number, "$")[1])

	if err != nil {
		return 0, errors.New("Error occured. Please write an valid expression")
	}

	valFromMemory, err := kv.Get(key)

	if err != nil {
		return 0, errors.New("Error occured. Key not found in memory")
	}

	return valFromMemory, nil
}
