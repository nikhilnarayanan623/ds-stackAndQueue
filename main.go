package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/stackAndQueue/queue"
	"github.com/nikhilnarayanan623/stackAndQueue/stack"
)

func main() {
	runQueue()
	runStack()
}

func runQueue() {
	fmt.Println("Queue")
	q := queue.NewQueue()

	q.MultipleEnqueue()
	q.DisplayOrder()
}

func runStack() {
	fmt.Println("Stack")

	s := stack.NewStack()
	s.DisplayOrder()

	s.MultiplePush()
	s.DisplayOrder()

	s.Pop()
	s.DisplayOrder()
}
