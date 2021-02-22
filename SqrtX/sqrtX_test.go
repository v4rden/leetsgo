package SqrtX

import (
	"fmt"
	"testing"
)

func TestSqrtXTableDriven(t *testing.T) {
	var tests = []struct {
		input  int
		result int
	}{
		{0, 0},
		{4, 2},
		{8, 2},
		{9, 3},
		{10, 3},
		{16, 4},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d", test.input)
		t.Run(testName,
			func(t *testing.T) {
				ans := mySqrt(test.input)
				if ans != test.result {
					t.Errorf("got %d wan't %d", ans, test.result)
				}
			})
	}
}
