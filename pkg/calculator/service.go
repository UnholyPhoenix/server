package calculator

import (
	"strconv"
	"strings"
)

// Calculate value from expression in string
func Calculate(expression string) string {
	// Splits input by space
	statment := strings.Fields(expression)

	if len(statment) != 3 {
		return "Error occured. Please write an valid expression.\n"
	}

	// Expression parts
	firstNumber := statment[0]
	operator := statment[1]
	secondNumber := statment[2]

	firstN, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		return "Error occured. Please write an valid expression.\n"
	}

	secondN, err := strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		return "Error occured. Please write an valid expression.\n"
	}

	switch {
	case operator == "+":
		result := firstN + secondN
		return strconv.FormatFloat(result, 'f', 2, 64)
	case operator == "-":
		return strconv.FormatFloat(firstN-secondN, 'f', 2, 64)
	case operator == "*":
		return strconv.FormatFloat(firstN*secondN, 'f', 2, 64)
	case operator == "/" && secondN == 0.00:
		return "\nNot possible"
	case operator == "/":
		return strconv.FormatFloat(firstN/secondN, 'f', 2, 64)
	default:
		return "\nInvalid operator. Allowed operators are '+', '-', '*' and '/'"
	}
}
