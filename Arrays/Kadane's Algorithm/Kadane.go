package main

import "fmt"

//import "fmt"

func sumArray(arr []int) int {
	arraySum := 0
	for i := 0; i < len(arr); i++ {
		arraySum += arr[i]
	}
	return arraySum
}

func kadane(arr []int) int {
	sum := arr[0]
	subArr := []int{}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {

			if i == j {
				subArr = []int{arr[i]}
				if sum < sumArray(subArr) {
					sum = sumArray(subArr)
				}
			} else {
				subArr = arr[i:j]
				if sumArray(subArr) > sum {
					//fmt.Println(subArr)

					sum = sumArray(subArr)
					//fmt.Println(sum)
				}
			}
		}
	}
	return sum
}

func main() {
	a := []int{-2, 1, 3, -4, -1, -2, -1, -5, -4}
	fmt.Println(kadane(a))
}
