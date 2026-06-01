func maxProfit(prices []int) int {
	
	maxProfit := 0
	maxPrice := prices[len(prices)-1]
	
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i+1] > maxPrice {
			maxPrice = prices[i+1]
		}
		profit := maxPrice - prices[i]
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}
