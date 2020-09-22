/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
						 Not 7-1 = 6, as selling price needs to be larger than buying price.
						 
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{7,1,5,3,6,4}, {7,6,4,3,1}}
	for _, test := range tests {
		log.Printf("maxProfit(%v) = %d\n", test, maxProfit(test))
	}
}

func maxProfit(prices []int) int {
		// at each element, we need to know the lowest element seen before that.
		// if bought at that price and sold at current price, that's the profit.
		var lowestSoFar int
		highestProfit := 0

		for i := 0; i < len(prices); i++ {
			if i == 0 {
				// initiate lowest
				lowestSoFar = prices[i]

			} else {
				profit := prices[i]-lowestSoFar
				if profit > highestProfit {
					highestProfit = profit
				}

				if prices[i] < lowestSoFar {
					lowestSoFar = prices[i]
				}
			}
		}

		return highestProfit
}
