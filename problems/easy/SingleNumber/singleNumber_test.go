package SingleNumber

import "testing"

func TestSingleNumberTableDriven(t *testing.T) {
	tests := []struct {
		id     string
		input  []int
		result int
	}{
		{"T1", []int{1, 2, 3, 4, 1, 2, 4}, 3},
		{"T2", []int{1}, 1},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := singleNumber(test.input)
			if ans != test.result {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}
