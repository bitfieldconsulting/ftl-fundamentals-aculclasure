// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

var errDivideByZero = errors.New("cannot divide by 0")

// Add accepts at least 2 addends and a variable number of extra
// addends and returns the result of adding them all together.
func Add(a, b float64, extra ...float64) float64 {
	sum := a + b
	for _, v := range extra {
		sum += v
	}
	return sum
}

// Subtract accepts at least 2 arguments and a variable number of
// extra arguments and returns the result of subtracting them in the
// order they are given.
func Subtract(a, b float64, extra ...float64) float64 {
	diff := a - b
	for _, v := range extra {
		diff -= v
	}
	return diff
}

// Multiply accepts 2 numbers and a variable number of extra numbers
// and returns their product.
func Multiply(a, b float64, extra ...float64) float64 {
	product := a * b
	for _, v := range extra {
		product *= v
	}
	return product
}

// Divide accepts 2 numbers and a variable number of extra numbers
// and returns the quotient of dividing them in the order they are
// given. If a divide-by-zero condition is encountered, then an
// error is returned.
func Divide(a, b float64, extra ...float64) (float64, error) {
	if b == 0 {
		return 0, errDivideByZero
	}
	quotient := a / b
	for _, v := range extra {
		if v == 0 {
			return 0, errDivideByZero
		}
		quotient /= v
	}
	return quotient, nil
}

// Sqrt takes a positive number and returns its square root. If a
// negative number is given, then an error is returned.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errDivideByZero
	}
	return math.Sqrt(a), nil
}

// Evaluate accepts a binary mathematical expression represented as a string
// and returns the evaluated value of the expression. If the expression does
// not contain a valid mathematical operator or represents an invalid operation
// (like dividing by 0), then an error is returned.
func Evaluate(expression string) (float64, error) {
	var (
		a, b     float64
		operator string
	)
	_, err := fmt.Sscanf(expression, "%f%s%f", &a, &operator, &b)
	if err != nil {
		return 0, err
	}
	switch operator {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("want a valid operator (+, -, *, /), got %s", operator)
	}
}
