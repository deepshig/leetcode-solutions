package main

import (
	"fmt"
	"strconv"
)

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
		fmt.Printf(" %d ->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
	return
}

func getDecimalValue(head *Node) int {
	binaryNumberStr := ""
	curr := head
	for curr != nil {
		binaryNumberStr += strconv.Itoa(curr.Val)
		curr = curr.Next
	}

	numberInt, err := strconv.ParseInt(binaryNumberStr, 2, 64)
	if err != nil {
		return 0
	}
	return int(numberInt)
}

func main() {

	fmt.Println("Solution 1")
	head := &Node{Val: 1}
	head.push(0)
	head.push(1)
	head.print()
	fmt.Println("Number : ", getDecimalValue(head))

	fmt.Println("Solution 2")
	head = &Node{Val: 0}
	head.print()
	fmt.Println("Number : ", getDecimalValue(head))

	fmt.Println("Solution 3")
	head = &Node{Val: 1}
	head.push(0)
	head.push(0)
	head.push(1)
	head.push(0)
	head.push(0)
	head.push(1)
	head.push(1)
	head.push(1)
	head.push(0)
	head.push(0)
	head.push(0)
	head.push(0)
	head.push(0)
	head.push(0)
	head.print()
	fmt.Println("Number : ", getDecimalValue(head))

	fmt.Println("Solution 4")
	head = &Node{Val: 1}
	head.print()
	fmt.Println("Number : ", getDecimalValue(head))

	fmt.Println("Solution 5")
	head = &Node{Val: 0}
	head.push(0)
	head.print()
	fmt.Println("Number : ", getDecimalValue(head))
}
