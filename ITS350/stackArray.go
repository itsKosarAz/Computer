package main

import (
	"fmt"
)

var array [5]int
var counter int

func main() {
	input()
	print()
}

func input() {
	for i := 0; i < len(array); i++ {
		fmt.Println("Enter a number")
		var tmp int
		fmt.Scanln(&tmp)
		push(tmp)
	}
}

func push(x int) {
	if counter == len(array) {
		fmt.Println("Stack is full")
		return
	}
	array[counter] = x
	counter++
}

func pop() int {
	counter--
	tmp := array[counter]
	array[counter] = 0
	return tmp
}

func print() {
	/*for i := 0; i < counter; i++ {
		fmt.Println(array[i])
	}*/
	fmt.Println(array)
}

func popAll() {
	tmp := counter
	for i := 0; i < tmp; i++ {
		pop()
	}
}
