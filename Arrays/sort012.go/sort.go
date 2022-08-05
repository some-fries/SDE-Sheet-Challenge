package main

import "fmt"

func sorter(arr []int) []int {
	zeroArray := []int{}
	oneArray := []int{}
	twoArray := []int{}
	sortedArray := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroArray = append(zeroArray, arr[i])
		} else if arr[i] == 1 {
			oneArray = append(oneArray, arr[i])
		} else if arr[i] == 2 {
			twoArray = append(twoArray, arr[i])
		}
	}
	for _, val := range zeroArray {
		sortedArray = append(sortedArray, val)
	}
	for _, val := range oneArray {
		sortedArray = append(sortedArray, val)
	}
	for _, val := range twoArray {
		sortedArray = append(sortedArray, val)
	}

	return sortedArray
}

func main() {
	arr1 := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(sorter(arr1))
}
