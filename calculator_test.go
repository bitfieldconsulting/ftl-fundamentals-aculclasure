package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		a, b  float64
		extra []float64
		want  float64
	}{
		{
			name:  "Sum of 2 operands to give a positive number",
			a:     1,
			b:     1,
			extra: nil,
			want:  2,
		},
		{
			name:  "Sum of 3 positive numbers to give a positive number",
			a:     1,
			b:     1,
			extra: []float64{1},
			want:  3,
		},
		{
			name:  "Sum of 5 negative numbers to give a negative number",
			a:     -2,
			b:     -2,
			extra: []float64{-2, -2, -2},
			want:  -10,
		},
		{
			name:  "Sum of 4 numbers to give 0 sum",
			a:     -1,
			b:     -1,
			extra: []float64{1, 1},
			want:  0,
		},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b, tc.extra...)
		if tc.want != got {
			t.Errorf("%s: Add(%v, %v, %+v) want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		a, b  float64
		extra []float64
		want  float64
	}{
		{
			name:  "Difference of 2 equal positive numbers to give 0",
			a:     1,
			b:     1,
			extra: nil,
			want:  0,
		},
		{
			name:  "Difference of 2 equal negative numbers to give 0",
			a:     -1,
			b:     -1,
			extra: nil,
			want:  0,
		},
		{
			name:  "Difference of more than 2 arguments",
			a:     10,
			b:     1,
			extra: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:  0,
		},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b, tc.extra...)
		if tc.want != got {
			t.Errorf("%s: Subtract(%v, %v, %+v)) want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	numTestCases := 100
	maxInputValue := 10
	maxNumExtraInputsPerTestCase := 10
	type testCase struct {
		a, b  float64
		extra []float64
		want  float64
	}
	randomTestCases := make(chan testCase)
	go func() {
		for i := 0; i < numTestCases; i++ {
			a, b := float64(rand.Intn(maxInputValue)), float64(rand.Intn(maxInputValue))
			want := a * b
			numInputs := rand.Intn(maxNumExtraInputsPerTestCase)
			extra := []float64{}
			for j := 0; j < numInputs; j++ {
				randomInput := float64(rand.Intn(maxInputValue) + 1)
				want *= randomInput
				extra = append(extra, randomInput)
			}
			randomTestCases <- testCase{
				a:     a,
				b:     b,
				extra: extra,
				want:  want,
			}
		}
		close(randomTestCases)
	}()
	for tc := range randomTestCases {
		got := calculator.Multiply(tc.a, tc.b, tc.extra...)
		if tc.want != got {
			t.Errorf("Multiply(%v, %v, %+v) want %f, got %f",
				tc.a, tc.a, tc.extra, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, b        float64
		extra       []float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Division of 2 equal numbers to give 1",
			a:           100,
			b:           100,
			extra:       nil,
			want:        1,
			errExpected: false,
		},
		{
			name:        "Division of positive and negative numbers to give a negative number",
			a:           -10,
			b:           5,
			extra:       nil,
			want:        -2,
			errExpected: false,
		},
		{
			name:        "Division of 2 negative numbers to give a positive number",
			a:           -10,
			b:           -10,
			extra:       nil,
			want:        1,
			errExpected: false,
		},
		{
			name:        "Division with more than 2 arguments",
			a:           100,
			b:           2,
			extra:       []float64{2, 5},
			want:        5,
			errExpected: false,
		},
		{
			name:        "Division of 2 numbers to give divide-by-zero error",
			a:           1,
			b:           0,
			want:        0,
			extra:       nil,
			errExpected: true,
		},
		{
			name:        "Division of more than 2 numbers to give divide-by-zero error",
			a:           200,
			b:           2,
			extra:       []float64{2, 2, 0, 1},
			want:        0,
			errExpected: true,
		},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b, tc.extra...)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s: Divide(%v, %v, %+v): unexpected error status: %v",
				tc.name, tc.a, tc.b, tc.extra, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s: Divide(%v, %v, %+v) want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, want     float64
		errExpected bool
	}{
		{name: "Square root of evenly divisible squarable positive number", a: 4, want: 2, errExpected: false},
		{name: "Square root of positive decimal number", a: 0.25, want: 0.5, errExpected: false},
		{name: "Square root of negative number to give error", a: -4, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s: Sqrt(%f): unexpected error status: %v", tc.name, tc.a, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s: Sqrt(%f) want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name, expression string
		want             float64
		errExpected      bool
	}{
		{
			name:        "Expression with no operator to return an error",
			expression:  "2 plus 2",
			want:        0,
			errExpected: true,
		},
		{
			name:        "Expression with missing operand to return an error",
			expression:  "2 +",
			want:        0,
			errExpected: true,
		},
		{
			name:        "Expression with only 1 valid operand to return an error",
			expression:  "2 - not_a_number",
			want:        0,
			errExpected: true,
		},
		{
			name:        "Expression with only 1 valid operand to return an error",
			expression:  "1 one wun + 1",
			want:        0,
			errExpected: true,
		},
		{
			name:        "Addition of 2 positive numbers to give positive number",
			expression:  "1 + 1",
			want:        2,
			errExpected: false,
		},
		{
			name:        "Subtraction of 2 identical numbers to give 0",
			expression:  " 1 - 1 ",
			want:        0,
			errExpected: false,
		},
		{
			name:        "Multiplication of 2 positive numbers to give a positive number",
			expression:  "  2  *    2    ",
			want:        4,
			errExpected: false,
		},
		{
			name:        "Division of evenly divisible positive numbers to give a positive number",
			expression:  "10/2",
			want:        5,
			errExpected: false,
		},
		{
			name:        "Division by 0 to give an error",
			expression:  "10          /0",
			want:        0,
			errExpected: true,
		},
	}
	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expression)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s: (%s): unexpected error status: %v", tc.name, tc.expression, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s: (%s), want: %f, got %f", tc.name, tc.expression, tc.want, got)
		}
	}
}
