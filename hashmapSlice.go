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
	slice []*LinkedList
}

func main() {

}

func (h *Hashmap) initialize(n int) {
	h.slice = h.slice[:n]
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

func hashing(name string) int {
	sum := 0
	for _, x := range name {
		sum += int(x)
	}
	return sum
}

func (h *Hashmap) addH(name string) {
	index := hashing(name) % len(h.slice)
	h.slice[index].addLL(name)
}

func (h *Hashmap) removeH(name string) {
	index := hashing(name) % len(h.slice)
	h.slice[index].removeLL(name)
}

func (h *Hashmap) search(name string) {
	index := hashing(name) % len(h.slice)
	h.slice[index].searchLL(name)
}
