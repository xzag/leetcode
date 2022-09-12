package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		customerWealth := 0
		for _, balance := range account {
			customerWealth += balance
		}

		if customerWealth > max {
			max = customerWealth
		}
	}

	return max
}
