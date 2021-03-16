package BestTimeToBuyAndSellStock2

import "testing"

func TestBestTimeToBuyAndSellStock2TableDriven(t *testing.T) {
	tests := []struct {
		id     string
		input  []int
		result int
	}{
		{"T1",
			[]int{7, 1, 5, 3, 6, 4},
			7},
		{"T2",
			[]int{1, 5, 1, 5, 1},
			8},
		{"T3",
			[]int{1, 2, 3, 4, 5},
			4},
		{"T4",
			[]int{9001},
			0},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := maxProfit(test.input)
			if ans != test.result {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}
