package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

// testCase represents a case for testing non-error producing Calculator functions.
type testCase struct {
	a, b, want float64
	name       string
}

// errorTestCase represents a case for testing Calculator functions that return errors.
type errorTestCase struct {
	a, b, want  float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []*testCase{
		{a: 1, b: 1, want: 2, name: "Two positive numbers which sum to a positive number"},
		{a: -1, b: -1, want: -2, name: "Two negative numbers which sum to a negative number"},
		{a: -1, b: 1, want: 0, name: "Positive and negative number which sum to 0"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f) want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a * b
		got := calculator.Multiply(a, b)
		if want != got {
			t.Errorf("Multiply(%f, %f) want %f, got %f", a, b, want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []*errorTestCase{
		{a: 4, b: 2, want: 2, errExpected: false, name: "Division of evenly divisible positive number to give a positive quotient"},
		{a: -4, b: -2, want: 2, errExpected: false, name: "Division of 2 negative numbers (evenly divisible) to give a positive quotient"},
		{a: -4, b: 2, want: -2, errExpected: false, name: "Division of negative number by even number to give negative quotient"},
		{a: 4, b: 0, want: 0, errExpected: true, name: "Division by zero to return an error"},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if tc.errExpected && err == nil {
			t.Fatalf("%s: Divide(%f, %f) expected an error to be returned but got nil", tc.name, tc.a, tc.b)
		}

		if !tc.errExpected && err != nil {
			t.Fatalf("%s: Divide(%f, %f) returned an unexpected error: %v", tc.name, tc.a, tc.b, err)
		}

		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s: Divide(%f, %f) want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
