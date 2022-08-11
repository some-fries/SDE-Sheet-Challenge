package main

import "fmt"

func main() {
	stocks := []int{7, 6, 4, 3, 1}
	stockProfit := profit(stocks)
	fmt.Println(stockProfit)
}

func profit(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	var buyPrice, sellPrice int
	profit := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			buyPrice = arr[i]
			sellPrice = arr[j]
			if sellPrice-buyPrice > profit {
				profit = sellPrice - buyPrice
			}
		}
	}
	return profit
}
