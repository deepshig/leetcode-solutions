package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) push(val int) *Node {
	newNode := &Node{
		Val:  val,
		Next: nil,
	}
	n.Next = newNode
	n = newNode
	return n
}

func arrayToLinkedList(arr []int, cyclePos int) *Node {
	head := &Node{
		Val:  arr[0],
		Next: nil,
	}
	cycleStart, tail := head, head

	for _, v := range arr[1:] {
		tail = tail.push(v)
	}

	if cyclePos >= 0 {
		for i := 0; i < cyclePos; i++ {
			cycleStart = cycleStart.Next
		}
		tail.Next = cycleStart
	}

	// cur := head
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println("i = ", i, "val = ", cur.Val)
	// 	cur = cur.Next
	// }
	// fmt.Println("total = ", len(arr))

	return head
}

func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {
	list := arrayToLinkedList([]int{3, 2, 0, -4}, 1)
	fmt.Println("Solution 1 : ", hasCycle(list))

	list = arrayToLinkedList([]int{1, 2}, 0)
	fmt.Println("Solution 2 : ", hasCycle(list))

	list = arrayToLinkedList([]int{1}, -1)
	fmt.Println("Solution 3 : ", hasCycle(list))
}
