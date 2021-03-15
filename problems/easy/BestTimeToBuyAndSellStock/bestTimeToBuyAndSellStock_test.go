package BestTimeToBuyAndSellStock

import "testing"

func TestBestTimeToBuyAndSellStockTableDriven(t *testing.T) {
	tests := []struct {
		id     string
		input  []int
		result int
	}{
		{"T1",
			[]int{7, 1, 5, 3, 6, 4},
			5},
		{"T2",
			[]int{7, 6, 5, 3, 2},
			0},
		{"T3",
			[]int{7, 8, 24, 1, 6, 15},
			17},
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
