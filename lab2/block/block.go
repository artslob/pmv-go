package block

import "fmt"

type Block interface {
	fmt.Stringer
	GetAddress() Block
	Append(Block)
	GetNext() Block
	GetId() int
}

type SimpleBlock struct {
	Id     int
	Text   string
	Next   Block
	Branch Block
	Level  int
}

func (b *SimpleBlock) GetAddress() Block {
	return b
}

func (b *SimpleBlock) Append(next Block) {
	b.Next = next
}

func (b *SimpleBlock) GetNext() Block {
	return b.Next
}

func (b *SimpleBlock) GetId() int {
	return b.Id
}

func (b *SimpleBlock) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.Id, b.Text)
}

/*** BLOCK STACK ***/

type BlockStack []Block

func (s *BlockStack) Push(b Block) {
	*s = append(*s, b)
}

func (s *BlockStack) Pop() Block {
	if len(*s) == 0 {
		panic("empty stack")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *BlockStack) Peek() Block {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *BlockStack) Size() int {
	return len(*s)
}

func (s *BlockStack) Empty() bool {
	return s.Size() == 0
}
