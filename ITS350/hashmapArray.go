package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	name string
}

type Hashmap struct {
	array [10]*LinkedList
}

func main() {

}

func (l *LinkedList) addLL(name string) {
	if l.head == nil {
		l.head = &Node{nil, name}
	} else {
		tmp := l.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &Node{nil, name}
	}
}

func (l *LinkedList) removeLL(name string) {
	if l.head.name == name {
		l.head = l.head.next
	} else {
		tmp := l.head
		for tmp.next.name != name {
			if tmp.next == nil {
				fmt.Println("Name does not exist")
				return
			}
			tmp = tmp.next
		}
		tmp.next = tmp.next.next
	}
}

func (l *LinkedList) searchLL(name string) {
	if l.head.name == name {
		fmt.Println("Name found")
	} else {
		tmp := l.head
		for tmp.next.name != name {
			if tmp.next == nil {
				fmt.Println("Name not found")
				return
			}
			tmp = tmp.next
		}
		fmt.Println("Name found")
	}
}

func hashing(Name string) int {
	sum := 0
	for _, x := range Name {
		sum += int(x)
	}
	return sum % 10
}

func (h *Hashmap) addH(Name string) {
	index := hashing(Name)
	h.array[index].addLL(Name)
}

func (h *Hashmap) removeH(Name string) {
	index := hashing(Name)
	h.array[index].removeLL(Name)
}

func (h *Hashmap) search(Name string) {
	index := hashing(Name)
	h.array[index].searchLL(Name)
}
