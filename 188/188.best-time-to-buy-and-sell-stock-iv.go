package leetcode

/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	l := len(prices)
	if k >= l/2 {
		return quickSolve(prices)
	}

	profits := make([][]int, k+1, k+1)
	for i := range profits {
		profits[i] = make([]int, l, l)
	}

	for i := 1; i <= k; i++ {
		profit := -prices[0]
		// Search what day(j) it should sell & buy a stock
		for j := 1; j < l; j++ {
			// Sell a stock
			if profits[i][j-1] < prices[j]+profit {
				profits[i][j] = prices[j] + profit
			} else {
				profits[i][j] = profits[i][j-1]
			}
			// Buy a stock
			if profit < profits[i-1][j-1]-prices[j] {
				profit = profits[i-1][j-1] - prices[j]
			}
		}
	}

	return profits[k][l-1]
}

func quickSolve(prices []int) int {
	p := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			p += prices[i] - prices[i-1]
		}
	}
	return p
}

// @lc code=end
