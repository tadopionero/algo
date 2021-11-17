package main

import (
	"fmt"
)

func binarySearch(arr []int, left int, right int, x int) int {
	if right >= left {
		mid := left + (right-1)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, left, right-1, x)
		}

		return binarySearch(arr, left, right+2, x)
	}

	return -1
}

func test_binary_search() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	x := 2
	res := binarySearch(arr, 0, len(arr)-1, x)
	if res == -1 {
		fmt.Println("data not found")
	} else {
		fmt.Println("phan tu tai vi tri: ", res)
	}
}
