package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

func TestAddSubtractMultiply(t *testing.T) {
	testCases := []struct {
		name                       string
		a, b                       float64
		extra                      []float64
		wantSum, wantDiff, wantPrd float64
	}{
		{
			name:     "Positive operands",
			a:        3,
			b:        2,
			wantSum:  5,
			wantDiff: 1,
			wantPrd:  6,
		},
		{
			name:     "Negative operands",
			a:        -2,
			b:        -2,
			wantSum:  -4,
			wantDiff: 0,
			wantPrd:  4,
		},
		{
			name:     "Mix of positive and negative operands",
			a:        -5,
			b:        5,
			wantSum:  0,
			wantDiff: -10,
			wantPrd:  -25,
		},
		{
			name:     "Mix of positive and negative operands",
			a:        5,
			b:        -5,
			wantSum:  0,
			wantDiff: 10,
			wantPrd:  -25,
		},
		{
			name:     "Zero for both operands",
			a:        0,
			b:        0,
			wantSum:  0,
			wantDiff: 0,
			wantPrd:  0,
		},
		{
			name:     "More than 2 operands (all positive)",
			a:        8,
			b:        1,
			extra:    []float64{3, 2},
			wantSum:  14,
			wantDiff: 2,
			wantPrd:  48,
		},
		{
			name:     "More than 2 operands (all negative)",
			a:        -6,
			b:        -3,
			extra:    []float64{-2, -1},
			wantSum:  -12,
			wantDiff: 0,
			wantPrd:  36,
		},
		{
			name:     "More than 2 operands (mixed negative and positive)",
			a:        4,
			b:        -3,
			extra:    []float64{2},
			wantSum:  3,
			wantDiff: 5,
			wantPrd:  -24,
		},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b, tc.extra...)
		if tc.wantSum != got {
			t.Errorf("%s: Add(%f, %f, %+v): want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.wantSum, got)
		}
		got = calculator.Subtract(tc.a, tc.b, tc.extra...)
		if tc.wantDiff != got {
			t.Errorf("%s: Subtract(%f, %f, %+v): want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.wantDiff, got)
		}
		got = calculator.Multiply(tc.a, tc.b, tc.extra...)
		if tc.wantPrd != got {
			t.Errorf("%s: Multiply(%f, %f, %+v): want %f, got %f",
				tc.name, tc.a, tc.b, tc.extra, tc.wantPrd, got)
		}
	}
}

func TestMultiplyWithRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a, b := float64(rand.Intn(10)), float64(rand.Intn(10))
		want := a * b
		numInputs := rand.Intn(100)
		extra := []float64{}
		for j := 0; j < numInputs; j++ {
			randomInput := float64(rand.Intn(10) + 1)
			want *= randomInput
			extra = append(extra, randomInput)
		}
		got := calculator.Multiply(a, b, extra...)
		if want != got {
			t.Errorf("Multiply(%f, %f, %+v): want: %f, got %f",
				a, b, extra, want, got)
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
			name:        "Expression with no spacing between operands and operator",
			expression:  "10/2",
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
			expression:  "10 / 2",
			want:        5,
			errExpected: false,
		},
		{
			name:        "Division by 0 to give an error",
			expression:  "10          / 0",
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
