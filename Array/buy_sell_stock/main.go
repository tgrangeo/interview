package main

import (
	"fmt"
	"slices"
)

func oneTimeProfit(prices []int) (int, int, int) {
	min := slices.Min(prices)
	max := 0
	for i := slices.Index(prices, min); i < len(prices); i++ {
		if prices[i] > max{
			max = prices[i]
		}
	}
	return slices.Index(prices, min),slices.Index(prices, max), max - min
}

func maxProfit(prices []int) int {
	maxProfit := 0
	for i,p := range prices{
		if (i > 0 && p > prices[i-1]){
			maxProfit += p - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	dayToBuy, dayToSell, profit := oneTimeProfit(prices)
	fmt.Println(dayToBuy,dayToSell, profit)
	fmt.Println(maxProfit(prices))
}
