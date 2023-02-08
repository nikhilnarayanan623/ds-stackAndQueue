package stack

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-stackAndQueue/stack/interfaces"
)

type node struct {
	data int
	next *node
}

// stack folow LIFO method Last In First Out
type stack struct {
	head *node
}

func NewStack() interfaces.StackInterface {
	return &stack{}
}

// function to insert a new value in to stack
func (s *stack) Push(data int) {

	newNode := &node{data: data}

	//push head to the next of newNode's next if head is nil or not
	newNode.next = s.head
	//change newNode as head
	s.head = newNode
}

// funciton to push multiple values to stack from user value
func (s *stack) MultiplePush() {
	var limit, data int

	fmt.Print("Enter how much values you need to enter: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter value: ")
		fmt.Scan(&data)
		s.Push(data)
	}
}

// funciton to delete a value from stack
// in stack head value been deleted
func (s *stack) Pop() {
	if checkStackEmpty(s.head, true) {
		return
	}
	//move head to its next
	s.head = s.head.next
}

// function to display all values in tha stack
func (s *stack) DisplayOrder() {

	if checkStackEmpty(s.head, true) {
		return
	}

	currentNode := s.head

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

// function to check the stack is empty
func checkStackEmpty(node *node, show bool) bool {

	if node == nil && show {
		fmt.Println("The stack is empty")
	}
	return node == nil
}
