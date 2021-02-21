package AddBinary

import (
	"fmt"
	"testing"
)

func TestAddBinaryTableDriven(t *testing.T) {
	var tests = []struct {
		input  []string
		result string
	}{
		{[]string{"11", "1"}, "100"},
		{[]string{"1010", "1011"}, "10101"},
		{[]string{"0", "1011"}, "1011"},
		{[]string{"1", "0"}, "1"},
		{[]string{"111", "1010"}, "10001"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v", test.input)
		t.Run(testName, func(t *testing.T) {
			ans := addBinary(test.input[0], test.input[1])
			if ans != test.result {
				t.Errorf("got %s want %s", ans, test.result)
			}
		})
	}
}
