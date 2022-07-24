package main

import (
	"fmt"
	"math"
)

func numberOfCoinRecursive(coins []int, totalCoin int, quantity int) int {

	if quantity == 0 {
		return 0
	}
	var result = math.MaxInt
	for i := 0; i < totalCoin; i++ {
		if coins[i] <= quantity {
			var subResult = 1 + numberOfCoinRecursive(coins, totalCoin, quantity-coins[i])
			if subResult < result {
				result = subResult
			}
		}
	}
	return result
}
func main() {
	var coins = []int{1, 10, 25}
	fmt.Printf("You receive %d coins in total", numberOfCoinRecursive(coins, len(coins), 30))
}
