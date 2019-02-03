/*
Given an array of consecutive daily prices for a particular stock, return the maximum profit that can be made by buying and selling a unit of stock during that time period.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}}

	for _, test := range tests {
		fmt.Println(maxProfitLinear(test))
	}
}

func maxProfitLinear(prices []int) int {
	// keep track of minimum price seen so far while iterating through array, and see the profilt that can be made at current price if bought at that minimum
	minPrice := 0
	maxProf := 0

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			minPrice = prices[i]
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if (prices[i] > minPrice) && ((prices[i]-minPrice) > maxProf) {
			maxProf = prices[i]-minPrice
		}
	}

	return maxProf
}

func maxProfit(prices []int) int {
	maxProf := 0
	for i := 0; i < len(prices); i++ {
		buy := prices[i]

		for j := i+1; j < len(prices); j++ {
			sell := prices[j]
			if sell > buy {
				if (sell - buy) > maxProf {
					maxProf = sell - buy
				}
			}
		}
	}

	return maxProf
}
