package DP__Multidimensional

import "fmt"

// maxProfit calculates the maximum possible profit from buying and selling stocks with a given transaction fee.
// The function takes a slice of integers representing the stock prices and an integer representing the transaction fee.
// It returns the maximum possible profit.
func maxProfit(prices []int, fee int) int {
	// Handle the edge case where the prices slice is empty.
	if len(prices) == 0 {
		return 0
	}

	// Initialize the hold and sold variables to represent the maximum profit when holding or selling a stock.
	// Initially, we hold a stock with a price of -prices[0], meaning we have a "debt" of the initial stock price.
	// We have not sold any stock yet, so the sold variable is initialized to 0.
	hold, sold := -prices[0], 0

	// Iterate through the prices slice starting from the second price (index 1).
	for i := 1; i < len(prices); i++ {
		// Get the current price.
		price := prices[i]

		// Update the hold variable to be the maximum of the current hold value and the value of selling the current stock and buying a new one.
		// This represents the maximum profit when holding a stock after the current transaction.
		hold = max(hold, sold-price)

		// Update the sold variable to be the maximum of the current sold value and the value of selling the current stock and paying the transaction fee.
		// This represents the maximum profit when selling a stock after the current transaction.
		sold = max(sold, hold+price-fee)
	}

	// Return the maximum possible profit, which is the sold value after the last transaction.
	return sold
}

// max returns the maximum of two integers.
func max(a, b int) int {
	// If a is greater than b, return a; otherwise, return b.
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	maxProfit := maxProfit(prices, fee)
	fmt.Printf("Maximum possible profit: %d\n", maxProfit)
}
