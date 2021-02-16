package stack

import "fmt"

type Stack []int

func NewStack() Stack {
	return Stack{}
}

func (s Stack) Size() int {
	return len(s)
}

func (s *Stack) Push(el int) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (int, error) {
	size := s.Size()
	if size == 0 {
		return 0, fmt.Errorf("%s", "Stack is emtpy")
	}
	el := (*s)[size-1]
	*s = (*s)[:size-1]
	return el, nil
}