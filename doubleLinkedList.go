package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	prev  *Node
	next  *Node
	value int
}

func main() {

}

func (l *LinkedList) add(value int) {
	if l.head == nil {
		l.head = &Node{nil, nil, value}
		l.tail = l.head
	} else {
		tmp := l.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &Node{tmp, nil, value}
		l.tail = tmp.next
	}
}

func (l *LinkedList) remove(value int) {
	if l.head.value == value {
		l.head = l.head.next
	} else {
		tmp := l.head
		for tmp.next.value != value {
			if tmp.next == nil {
				fmt.Println("Value does not exist")
				return
			}
			tmp = tmp.next
		}
		if tmp.next.next != nil {
			tmp.next = tmp.next.next
			tmp.next.prev = tmp
		} else {
			tmp.next = nil
			l.tail = tmp
		}
	}
}

func (l *LinkedList) search(value int) {
	if l.head.value == value {
		fmt.Println("Value found")
	} else {
		tmp := l.head
		for tmp.next.value != value {
			if tmp.next == nil {
				fmt.Println("Value not found")
				return
			}
			tmp = tmp.next
		}
		fmt.Println("Value found")
	}
}
