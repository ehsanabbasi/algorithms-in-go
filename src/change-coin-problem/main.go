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
	fmt.Printf("%v\n", result)
}
func main() {
	var coins = []int{1, 10, 30}
	numberOfCoin(coins, 15)
}
