package main

import "fmt"

func bubbleSort(arr []int) {
	size := len(arr)

	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	fmt.Println(arr)
}

func test_bubble_sort() {
	arr := []int{1, 2, 5, 4, 6, 3}
	bubbleSort(arr)
}
