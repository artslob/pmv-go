package lab2

import "fmt"

type BlockClass int

const (
	DEFAULT BlockClass = iota
	START
	END
)

type Block struct {
	id     int
	text   string
	next   *Block
	branch *Block
	class  BlockClass
}

func (b *Block) String() string {
	label := ""
	switch b.class {
	case START:
		label = "START"
	case END:
		label = "END"
	default:
		label = b.text
	}
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.id, label)
}

type BlockStack []*Block

func (s *BlockStack) Push(b *Block) {
	*s = append(*s, b)
}

func (s *BlockStack) Pop() *Block {
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
