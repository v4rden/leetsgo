package ClimbingStairs

import (
	"fmt"
	"testing"
)

func TestClimbingStairsTableDriven(t *testing.T) {
	var tests = []struct {
		input  int
		result int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{45, 1836311903},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d", test.input)
		t.Run(testName, func(t *testing.T) {
			ans := climbStairs(test.input)
			if ans != test.result {
				t.Errorf("got %d want %d", ans, test.result)
			}
		})
	}
}
