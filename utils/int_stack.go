package utils

type IntStack []int

func (s *IntStack) Push(b int) {
	*s = append(*s, b)
}

func (s *IntStack) Pop() int {
	if len(*s) == 0 {
		panic("empty stack")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *IntStack) Peek() *int {
	if s.Empty() {
		return nil
	}
	return &(*s)[len(*s)-1]
}

func (s *IntStack) Size() int {
	return len(*s)
}

func (s *IntStack) Empty() bool {
	return s.Size() == 0
}
