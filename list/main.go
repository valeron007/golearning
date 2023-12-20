package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (list *Node) Sum() int {
	if list == nil {
		return 0
	}
	return list.data + list.next.Sum()
}

func (list *Node) add(value int) int {
	elem := &Node{value, nil}

	if list.next == nil {
		list.next = elem
		return value
	}

	current := list.next
	for current.next != nil {
		current = current.next
	}
	current.next = elem
	return value
}

func (list *Node) print() {
	current := list
	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("%d \n", current.data)
}

func main() {
	node := &Node{1, nil}
	node.add(2)
	node.add(5)
	node.print()
}
