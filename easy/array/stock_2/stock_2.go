package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	buyingPrice, sellingPrice, profit := 0, 0, 0
	stockBought := false

	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i] < prices[i-1] {
			if stockBought {
				sellingPrice = prices[i-1]
				profit += sellingPrice - buyingPrice
				stockBought = false
			}
			if !stockBought {
				buyingPrice = prices[i]
				stockBought = true
			}
		}

		if !stockBought && i < len(prices)-1 && prices[i] < prices[i+1] {
			buyingPrice = prices[i]
			stockBought = true
		}

		if stockBought && i == len(prices)-1 {
			sellingPrice = prices[i]
			profit += sellingPrice - buyingPrice
			stockBought = false
		}
	}
	return profit
}

func main() {
	fmt.Println("Solution 1 :", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Solution 2 :", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println("Solution 3 :", maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println("Solution 4 :", maxProfit([]int{1, 7, 2, 3, 6, 7, 6, 7}))
}
