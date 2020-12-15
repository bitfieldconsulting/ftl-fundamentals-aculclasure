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

// variadicTestCase represents a test case for a Calculator function that
// accepts a variable number of inputs.
type variadicTestCase struct {
	inputs []float64
	want   float64
	name   string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []*variadicTestCase{
		{inputs: []float64{1, 1, 1}, want: 3, name: "Sum of 3 positive numbers to give a positive number"},
		{inputs: []float64{-1, -1, -1}, want: -3, name: "Sum of 3 negative numbers to give a negative number"},
		{inputs: []float64{-1, 1, 0}, want: 0, name: "Sum of negative number and positive number"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("%s: Add(%v) want %f, got %f", tc.name, tc.inputs, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []*variadicTestCase{
		{inputs: []float64{}, want: 0, name: "Difference of empty slice of arguments"},
		{inputs: []float64{100}, want: 100, name: "Difference of single argument"},
		{inputs: []float64{1, 1, 1}, want: -1, name: "Difference of 3 positive numbers to give a negative number"},
		{inputs: []float64{5, 1, 1}, want: 3, name: "Difference of 3 positive numbers to give a positive number"},
		{inputs: []float64{-1, -1, -1, -1}, want: 2, name: "Difference of 4 negative numbers to give a positive number"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.inputs...)
		if tc.want != got {
			t.Errorf("%s: Subtract(%v) want %f, got %f", tc.name, tc.inputs, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	numTestCases := 100
	maxInputValue := 10
	maxNumInputsPerTestCase := 10
	testCases := make(chan *variadicTestCase)

	go func() {
		for i := 0; i < numTestCases; i++ {
			numInputs := rand.Intn(maxNumInputsPerTestCase)
			inputs := []float64{}
			var want float64 = 0
			if numInputs > 0 {
				want = 1
			}
			for j := 0; j < numInputs; j++ {
				randomInput := float64(rand.Intn(maxInputValue) + 1)
				want *= randomInput
				inputs = append(inputs, randomInput)
			}
			testCases <- &variadicTestCase{inputs: inputs, want: want}
		}
		close(testCases)
	}()

	for tc := range testCases {
		got := calculator.Multiply(tc.inputs...)
		if tc.want != got {
			t.Errorf("Multiply(%v) want %f, got %f", tc.inputs, tc.want, got)
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

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []*errorTestCase{
		{a: 4, want: 2, errExpected: false, name: "Square root of evenly squarable positive number"},
		{a: 0.25, want: 0.5, errExpected: false, name: "Square root of positive decimal number"},
		{a: -4, want: 0, errExpected: true, name: "Square root of negative number to produce an error"},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if tc.errExpected && err == nil {
			t.Fatalf("%s: Sqrt(%f) expected an error to be returned but got nil", tc.name, tc.a)
		}

		if !tc.errExpected && err != nil {
			t.Fatalf("%s: Sqrt(%f) returned an unexpected error: %v", tc.name, tc.a, err)
		}

		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s: Sqrt(%f) want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}
