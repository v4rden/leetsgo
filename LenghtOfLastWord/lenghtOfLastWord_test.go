package LenghtOfLastWord

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWordTableDriven(t *testing.T) {
	var tests = []struct {
		n      string
		result int
	}{
		{"Hello, world", 5},
		{" ", 0},
		{"   ", 0},
		{"no_spaces", 9},
		{"after ", 5},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := lengthOfLastWord(tt.n)
			if ans != tt.result {
				t.Errorf("got %d, want %d", ans, tt.result)
			}
		})
	}
}
