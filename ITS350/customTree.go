package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	left   *Node
	middle *Node
	right  *Node
	value  int
}

func main() {

}

func (n *Node) add(value int) {
	if n == nil {
		n = &Node{nil, nil, nil, value}
	} else if n.value > value {
		if n.left == nil {
			n.left = &Node{nil, nil, nil, value}
			return
		}
		n.left.add(value)
	} else if n.value < value {
		if n.right == nil {
			n.right = &Node{nil, nil, nil, value}
			return
		}
		n.right.add(value)
	} else {
		if n.middle == nil {
			n.middle = &Node{nil, nil, nil, value}
			return
		}
		n.middle.add(value)
	}
}

func (n *Node) search(value int) {
	if n == nil {
		fmt.Println("Value not found")
	} else if n.value > value {
		if n.left == nil {
			fmt.Println("Value not found")
			return
		}
		n.left.search(value)
	} else if n.value < value {
		if n.right == nil {
			fmt.Println("Value not found")
			return
		}
		n.right.search(value)
	} else {
		if n.middle == nil {
			fmt.Println("Value not found")
			return
		}
		n.middle.add(value)
	}
}
