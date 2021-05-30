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
		enqueue(tmp)
	}
}

func enqueue(x int) {
	if counter == len(array) {
		fmt.Println("Queue is full")
		return
	}
	array[counter] = x
	counter++
}

func dequeue() int {
	counter--
	tmp := array[0]
	for i := 0; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	array[counter] = 0
	return tmp
}

func print() {
	/*for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}*/
	fmt.Println(array)
}

func dequeueAll() {
	tmp := counter
	for i := 0; i < tmp; i++ {
		dequeue()
	}
}
