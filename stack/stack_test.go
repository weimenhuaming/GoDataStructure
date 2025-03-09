package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("栈内元素个数为%d\n", s.Len())
	s.Pop()
	fmt.Printf("栈内元素个数为%d\n", s.Len())
}
