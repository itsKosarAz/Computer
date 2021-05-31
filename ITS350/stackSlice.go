package main

import (
	"fmt"
)

var slice []int
var n int

func main() {
	input()
	pop()
	print()
}

func input() {
	fmt.Println("How many numbers?")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter a number")
		var tmp int
		fmt.Scanln(&tmp)
		push(tmp)
	}
}

func push(x int) {
	slice = append(slice, x)
}

func pop() int {
	lastIndex := len(slice) - 1
	tmp := slice[lastIndex]
	slice = slice[:lastIndex]
	return tmp
}

func print() {
	/*for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}*/
	fmt.Println(slice)
}

func popAll() {
	tmp := len(slice)
	for i := 0; i < tmp; i++ {
		pop()
	}
}
