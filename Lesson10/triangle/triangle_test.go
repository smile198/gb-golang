package triangle

import (
	"math"
	"testing"
)

func TestMain(t *testing.M) {
	t.Run()
}

func isCloseEnough(a, b float64) bool {
	accuracy := 0.0001

	return math.Abs(a-b) < accuracy
}

func TestTriangle_Square(t *testing.T) {
	cases := []struct {
		name           string
		input          Triangle
		expectedSquare float64
		expectedError  error
	}{
		{
			name: "Simple case",
			input: Triangle{
				SideA: 5,
				SideB: 6,
				SideC: 7,
			},
			expectedSquare: 14.696938,
		},
		{
			name: "Simple case 2",
			input: Triangle{
				SideA: 8,
				SideB: 9,
				SideC: 10,
			},
			expectedSquare: 34.197039,
		},
		{
			name: "Case with impossible triangle",
			input: Triangle{
				SideA: 1,
				SideB: 2,
				SideC: 3,
			},
			expectedError: ErrImpossibleTriangle,
		},
	}

	for _, cse := range cases {
		t.Run(cse.name, func(t *testing.T) {
			square, err := cse.input.Square()
			if err != nil {
				t.Fatal(err)
			}
			//if !errors.Is(err, ErrImpossibleTriangle) {
			//	t.Fatal(err)
			//}

			if !isCloseEnough(square, cse.expectedSquare) {
				t.Fatal("Err")
			}
		})
	}
}
