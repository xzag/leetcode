package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min, max int
	dp := make([][]int, len(prices))
	min = prices[0]
	dp[0] = []int{0, min}

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		dp[i] = []int{prices[i] - dp[i-1][1], min}
		if dp[i][0] > max {
			max = dp[i][0]
		}
	}

	return max
}
