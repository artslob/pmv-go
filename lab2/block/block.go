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
