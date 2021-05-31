package main

import (
	"fmt"
)

var slice []int
var n int

func main() {
	input()
	dequeueAll()
	print()
}

func input() {
	fmt.Println("How many numbers?")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter a number")
		var tmp int
		fmt.Scanln(&tmp)
		enqueue(tmp)
	}
}

func enqueue(x int) {
	slice = append(slice, x)
}

func dequeue() int {
	tmp := slice[0]
	slice = slice[1:]
	return tmp
}

func print() {
	/*for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}*/
	fmt.Println(slice)
}

func dequeueAll() {
	tmp := len(slice)
	for i := 0; i < tmp; i++ {
		dequeue()
	}
}
