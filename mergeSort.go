package main

import (
	"fmt"
)

func main() {
	slice := []int{4, 2, 6, 3, 8, 1, 5, 7, 0, 9}
	fmt.Println(mergeSort(slice))
}

func mergeSort(slice []int) []int {
	size := len(slice)
	if size == 1 {
		return slice
	}
	middle := len(slice) / 2
	left := slice[:middle]
	right := slice[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(sliceLeft []int, sliceRight []int) []int {
	slice := make([]int, len(sliceLeft)+len(sliceRight))
	x := 0
	y := 0
	z := 0
	for z < len(slice) {
		if x > len(sliceLeft)-1 {
			slice[z] = sliceRight[y]
			z++
			y++
		} else if y > len(sliceRight)-1 {
			slice[z] = sliceLeft[x]
			z++
			x++
		} else if sliceLeft[x] < sliceRight[y] {
			slice[z] = sliceLeft[x]
			z++
			x++
		} else {
			slice[z] = sliceRight[y]
			z++
			y++
		}
	}
	return slice
}
