// Package calculator provides a library for simple calculations in Go.
package calculator

import "fmt"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying
// them together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the
// first number by the second number. If the second number is 0,
// then an error is returned.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("got invalid value for b (%f), want any number other than 0", b)
	}

	return (a / b), nil
}
