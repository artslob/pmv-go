package block

type Stack []Block

func (s *Stack) Push(b Block) {
	*s = append(*s, b)
}

func (s *Stack) Pop() Block {
	if len(*s) == 0 {
		panic("empty stack")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Peek() Block {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}
