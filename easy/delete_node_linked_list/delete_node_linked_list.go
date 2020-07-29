package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) push(val int) {
	newNode := &Node{Val: val}
	curr := n
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	return
}

func (n *Node) print() {
	curr := n
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println()
}

func deleteNode(node *Node) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
	return
}

func main() {
	head := &Node{Val: 4}
	head.push(5)
	head.push(1)
	head.push(9)
	head.print()
	deleteNode(head.Next)
	head.print()
}
