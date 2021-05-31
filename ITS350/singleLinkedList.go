  
package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	next  *Node
	value int
}

func main() {

}

func (l *LinkedList) add(value int) {
	if l.head == nil {
		l.head = &Node{nil, value}
	} else {
		tmp := l.head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &Node{nil, value}
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
		tmp.next = tmp.next.next
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
