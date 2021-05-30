package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func main() {

}

func (n *Node) add(value int) {
	if n == nil {
		n = &Node{nil, nil, value}
	} else if n.value > value {
		if n.left == nil {
			n.left = &Node{nil, nil, value}
			return
		}
		n.left.add(value)
	} else if n.value < value {
		if n.right == nil {
			n.right = &Node{nil, nil, value}
			return
		}
		n.right.add(value)
	} else {
		fmt.Println("Value already exists")
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
		fmt.Println("Value found")
	}
}
