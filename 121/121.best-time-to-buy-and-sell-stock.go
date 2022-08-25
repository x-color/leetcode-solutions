package leetcode

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > profit {
			profit = price - min
		}
	}
	return profit
}

// @lc code=end
