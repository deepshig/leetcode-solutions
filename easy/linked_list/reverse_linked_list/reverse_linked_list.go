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
		fmt.Printf(" %d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
	return
}

func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next
	head.Next = nil

	for curr.Next != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	curr.Next = prev

	return curr
}

func main() {
	fmt.Println("Solution 1")
	head := &Node{Val: 1}
	head.push(2)
	head.push(3)
	head.push(4)
	head.push(5)
	head.print()
	reverse := reverseList(head)
	reverse.print()
}
