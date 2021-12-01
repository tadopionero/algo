package main

import "fmt"

func max(arr []int) int {
	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	fmt.Println(max)
	return max
}

func test_max() {
	arr := []int{1, 2, 5, 4, 6, 3}
	max(arr)
}
