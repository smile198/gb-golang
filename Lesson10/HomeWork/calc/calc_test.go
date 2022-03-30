package calc

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		name           string
		a              float32
		b              float32
		op             string
		expectedResult float32
		expectedError  error
	}{
		{
			name:           "1+2=3",
			a:              1,
			b:              2,
			op:             "+",
			expectedResult: 3,
		},
		{
			name:           "0-0=0",
			a:              0,
			b:              0,
			op:             "+",
			expectedResult: 0,
		},
		{
			name:           "1-(-2)=3",
			a:              1,
			b:              -2,
			op:             "-",
			expectedResult: 3,
		},
		{
			name:           "1*0=0",
			a:              1,
			b:              0,
			op:             "*",
			expectedResult: 0,
		},
		{
			name:           "3/2=1.5",
			a:              3,
			b:              2,
			op:             "/",
			expectedResult: 1.5,
		},
		{
			name:          "DivisionByZeroError",
			a:             3,
			b:             0,
			op:            "/",
			expectedError: DivisionByZeroError,
		},
		{
			name:           "3^2=9",
			a:              3,
			b:              2,
			op:             "^",
			expectedResult: 9,
		},
		{
			name:           "3!=6",
			a:              3,
			op:             "!",
			expectedResult: 6,
		},
		{
			name:           "0!=1",
			a:              0,
			op:             "!",
			expectedResult: 1,
		},
		{
			name:          "NegativeFactorialError",
			a:             -1,
			op:            "!",
			expectedError: NegativeFactorialError,
		},
		{
			name:          "UnknownOperatorError",
			a:             1,
			b:             2,
			op:            "log",
			expectedError: UnknownOperatorError,
		},
	}

	for _, cse := range cases {
		t.Run(cse.name, func(t *testing.T) {
			res, err := Calc(cse.a, cse.b, cse.op)

			if !errors.Is(err, cse.expectedError) {
				t.Fatal(err)
			}

			if res != cse.expectedResult {
				t.Fatalf("expecting: %f, got: %f", cse.expectedResult, res)
			}
		})
	}
}
