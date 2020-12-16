// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes a variable number of arguments and returns the
// result of adding them together.
func Add(inputs ...float64) float64 {
	var sum float64

	for _, v := range inputs {
		sum += v
	}
	return sum
}

// Subtract takes a variable number of arguments and returns the
// result of subtracting them in the order they are given.
func Subtract(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}

	if len(inputs) == 1 {
		return inputs[0]
	}

	difference := inputs[0]
	for _, v := range inputs[1:] {
		difference -= v
	}
	return difference
}

// Multiply takes a variable number of arguments and returns
// the result of multiplying them together.
func Multiply(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}

	var product float64 = 1
	for _, v := range inputs {
		product *= v
	}
	return product
}

// Divide a variable number of arguments and returns the result of
// dividing them by each other (starting from the first input). If
// a divide-by-zero happens at any point, then an error is returned.
func Divide(inputs ...float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, nil
	}

	if len(inputs) == 1 {
		return inputs[0], nil
	}

	quotient := inputs[0]
	for _, v := range inputs[1:] {
		if v == 0 {
			return 0, fmt.Errorf("got an invalid input value (%f), want any value other than 0", v)
		}
		quotient /= v
	}
	return quotient, nil
}

// Sqrt takes a positive number and returns its square root. If a
// negative number is given, then an error is returned.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("got invalid value for a (%f), want any number other than 0", a)
	}

	return math.Sqrt(a), nil
}

// Evaluate accepts a binary mathematical expression represented as a string
// and returns the evaluated value of the expression. If the expression does
// not contain a valid mathematical operator or represents an invalid operation
// (like dividing by 0), then an error is returned.
func Evaluate(expression string) (float64, error) {
	operatorIndex := strings.IndexFunc(expression, isOperator)
	if operatorIndex == -1 {
		return 0, fmt.Errorf("want an expression containing an operator (+, -, *, /), got %s", expression)
	}

	operator := rune(expression[operatorIndex])
	operands := strings.FieldsFunc(expression, func(r rune) bool { return r == operator })
	if len(operands) != 2 {
		return 0, fmt.Errorf("want 2 operands to work on (got %v)", operands)
	}

	operandValues := []float64{}
	for _, v := range operands {
		operand, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return 0, err
		}
		operandValues = append(operandValues, operand)
	}

	switch operator {
	case '+':
		return Add(operandValues...), nil
	case '-':
		return Subtract(operandValues...), nil
	case '*':
		return Multiply(operandValues...), nil
	default:
		return Divide(operandValues...)
	}
}

// isOperator accepts a rune r and returns true if r is a valid
// binary mathematical operator and false otherwise.
func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}
