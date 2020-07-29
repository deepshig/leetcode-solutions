package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) push(val int) {
	newNode := &Node{Val: val}
	head := n
	for head.Next != nil {
		head = head.Next
	}
	head.Next = newNode
	return
}

func (n *Node) print() {
	head := n
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
	return
}

func removeElements(head *Node, val int) *Node {
	if head == nil || (head.Next == nil && head.Val == val) {
		return nil
	}

	if head.Next == nil && head.Val != val {
		return head
	}

	prev, curr := head, head.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}

func main() {
	fmt.Println("Solution 1")
	head := &Node{Val: 1}
	head.push(2)
	head.push(6)
	head.push(3)
	head.push(4)
	head.push(5)
	head.push(6)
	head.print()
	head = removeElements(head, 6)
	head.print()

	fmt.Println("Solution 2")
	head = &Node{Val: 1}
	head.print()
	head = removeElements(head, 2)
	head.print()

	fmt.Println("Solution 3")
	head = &Node{Val: 1}
	head.print()
	head = removeElements(head, 1)
	head.print()

	fmt.Println("Solution 4")
	head = &Node{Val: 1}
	head.push(1)
	head.print()
	head = removeElements(head, 1)
	head.print()
}
