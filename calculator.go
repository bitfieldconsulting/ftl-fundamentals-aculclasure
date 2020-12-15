// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
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

// Sqrt takes a positive number and returns its square root. If a
// negative number is given, then an error is returned.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("got invalid value for a (%f), want any number other than 0", a)
	}

	return math.Sqrt(a), nil
}
