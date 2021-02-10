package CountAndSay

import (
	"fmt"
	"testing"
)

func TestCountAndSayTableDriven(t *testing.T) {
	var tests = []struct {
		n      int
		result string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := countAndSay(tt.n)
			if ans != tt.result {
				t.Errorf("got %s, want %s", ans, tt.result)
			}
		})
	}
}
