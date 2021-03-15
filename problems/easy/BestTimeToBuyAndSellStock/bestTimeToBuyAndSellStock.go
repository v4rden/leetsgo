package BestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	var maxDiff, currentDiff int

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		currentDiff = price - minPrice
		if currentDiff > maxDiff {
			maxDiff = currentDiff
		}
	}
	return maxDiff
}
