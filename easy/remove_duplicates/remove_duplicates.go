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

func deleteDuplicates(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next
	for curr != nil {
		if prev.Val == curr.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head
}

func main() {
	fmt.Println("Solution 1")
	head := &Node{Val: 1}
	head.push(1)
	head.push(2)
	head.print()
	noDuplicates := deleteDuplicates(head)
	noDuplicates.print()

	fmt.Println("Solution 2")
	head = &Node{Val: 1}
	head.push(1)
	head.push(2)
	head.push(3)
	head.push(3)
	head.print()
	noDuplicates = deleteDuplicates(head)
	noDuplicates.print()
}
