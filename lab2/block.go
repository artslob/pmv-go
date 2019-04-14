package lab2

import "fmt"

type Block struct {
	id     int
	text   string
	next   *Block
	branch *Block
	level  int
}

func (b *Block) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.id, b.text)
}

type BlockStack []*Block

func (s *BlockStack) Push(b *Block) {
	*s = append(*s, b)
}

func (s *BlockStack) Pop() *Block {
	if len(*s) == 0 {
		panic("empty stack")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *BlockStack) Peek() *Block {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *BlockStack) Size() int {
	return len(*s)
}
