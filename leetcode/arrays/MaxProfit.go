package arrays

// MaxProfit function that calculates the profit after buying and selling on any day.
func MaxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit = maxProfit + prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
