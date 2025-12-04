package utils

import "fmt"

type stack []int

// returns a new, empty stack
func NewStack() stack {
	return stack{}
}

// insert an element in the stack
func (s stack) Push(n int) stack {
	s = append(s, n)
	return s
}

// removes the top element of stack
func (s stack) Pop() (stack, int, error) {
	if len(s) == 0 {
		return s, -1, fmt.Errorf("stack is empty")
	}
	n := s[len(s)-1]
	s = s[0 : len(s)-1]
	return s, n, nil
}

// returns the top element of stack
func (s stack) Top() (int, error) {
	if len(s) == 0 {
		return -1, fmt.Errorf("stack is empty")
	}

	return s[len(s)-1], nil
}

// returns the number of element in stack
func (s stack) Len() int {
	return len(s)
}

func (s stack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}
