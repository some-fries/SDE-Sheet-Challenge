package main

import (
	"fmt"
	"sort"
)

func sortArray(arr []int) []int {
	newArr := arr
	sort.Ints(newArr)
	return newArr
}

/*
func perm(slice []int) [][]int {
	sortedSlice := sortArray(slice)
	twoDslice := [][]int{}
	newSlice := sortedSlice
	for i:=0; i<len(slice); i++{
		newSlice[0] = sortedSlice[i]
		for j:=1; j<len(slice); j++{
			newSlice[j] = sortedSlice[len(slice)-j]
		}
	}
	return [][]int{}
}
*/
func main() {
	a := []int{4, 2, 1, 10}
	//sort.Ints(a)
	fmt.Println(sortArray(a))
}
