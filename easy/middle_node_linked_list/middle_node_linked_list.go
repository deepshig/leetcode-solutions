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
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
	return
}

func middleNodeByCalculatingLength(head *Node) *Node {
	length, curr := 0, head
	for curr != nil {
		length++
		curr = curr.Next
	}
	midNode := (length / 2)
	curr = head
	for midNode > 0 {
		curr = curr.Next
		midNode--
	}
	return curr
}

func middleNode(head *Node) *Node {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func main() {
	fmt.Println("Solution 1")
	head := &Node{Val: 1}
	head.push(2)
	head.push(3)
	head.push(4)
	head.push(5)
	head.print()
	mid := middleNode(head)
	mid.print()

	fmt.Println("Solution 2")
	head = &Node{Val: 1}
	head.push(2)
	head.push(3)
	head.push(4)
	head.push(5)
	head.push(6)
	head.print()
	mid = middleNode(head)
	mid.print()
}
