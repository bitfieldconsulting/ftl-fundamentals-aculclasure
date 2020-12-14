package calculator_test

import (
	"calculator"
	"testing"
)

// testCase represents a case that can be tested in the various Test functions.
type testCase struct {
	a, b, want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []*testCase{
		{a: 1, b: 1, want: 2},
		{a: -1, b: -1, want: -2},
		{a: -1, b: 1, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
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
