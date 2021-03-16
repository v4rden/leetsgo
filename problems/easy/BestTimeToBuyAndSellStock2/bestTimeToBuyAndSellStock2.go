package BestTimeToBuyAndSellStock2

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var total int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}
