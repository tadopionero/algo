package main

import (
	"fmt"
)

func linearSearch(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 2, 4, 6, 5}

	fmt.Println(arr)
	x := 2
	res := linearSearch(arr, x)
	if res == -1 {
		fmt.Println("data not found")
	} else {
		fmt.Println("phan tu tai vi tri: ", res)
	}
}
