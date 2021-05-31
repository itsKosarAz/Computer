package main

import (
	"fmt"
)

func main() {
	array := [10]int{4, 2, 6, 3, 8, 1, 5, 7, 0, 9}
	fmt.Println(bubbleSort(array))
}

func bubbleSort(array [10]int) [10]int {
	boolean := false
	for i := len(array); i > 0; i-- {
		boolean = false
		for j := 0; j < i-1; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
				boolean = true
			}
		}
		if !boolean {
			return array
		}
	}
	return array
}
