package main

import (
	"fmt"
	"sort"
)

func numberOfCoin(coins []int, quantity int) {
	var coinSum int
	var coin int
	var i int
	var result []int

	if quantity < 0 {
		return
	}
	sort.Ints(coins)
	i = len(coins) - 1

	for i >= 0 && coinSum < quantity {
		coin = coins[i]
		for coin+coinSum <= quantity {
			result = append(result, coin)
			coinSum += coin
		}
		i--
	}
	fmt.Printf("You receive %d coins in total \n"+
		"Your coins in detail  %v\n"+
		"You asked for the amount of %d \n"+
		"Your coins input %v ", coinSum, result, quantity, coins)
}
func main() {
	var coins = []int{2, 10, 30}
	numberOfCoin(coins, 15)
}
