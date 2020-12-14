package calculator_test

import (
	"calculator"
	"testing"
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
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
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
	var want float64 = 6
	got := calculator.Multiply(3, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
