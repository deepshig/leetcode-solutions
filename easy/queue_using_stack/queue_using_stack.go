package main

import "fmt"

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{
		elements: []int{},
	}
}

func (s *Stack) Push(x int) {
	s.elements = append(s.elements, x)
	return
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return 0
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Peek() int {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}

type MyQueue struct {
	stack *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack: NewStack(),
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	temp := NewStack()
	for !q.stack.Empty() {
		element := q.stack.Pop()
		temp.Push(element)
	}

	q.stack.Push(x)
	for !temp.Empty() {
		element := temp.Pop()
		q.stack.Push(element)
	}

	return
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	return q.stack.Pop()
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}
	return q.stack.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.stack.Empty()
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	a := stack.Peek()
	b := stack.Pop()
	c := stack.Empty()
	d := stack.Peek()
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c, ", d = ", d)

	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	a = queue.Peek()
	b = queue.Pop()
	c = queue.Empty()
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
}
