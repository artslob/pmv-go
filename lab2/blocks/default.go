package blocks

import "fmt"

type DefaultBlock struct {
	Id            int
	Text          string
	FunctionCalls []string
	next          Block
}

func (b *DefaultBlock) GetAddress() Block {
	return b
}

func (b *DefaultBlock) SetNext(next Block) {
	b.next = next.GetAddress()
}

func (b *DefaultBlock) GetNext() Block {
	return b.next
}

func (b *DefaultBlock) SetBranch(Block) {
	panic("simple block do not provide branching")
}

func (b *DefaultBlock) GetBranch() Block {
	return nil
}

func (b *DefaultBlock) GetId() int {
	return b.Id
}

func (b *DefaultBlock) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.Id, b.Text)
}

func NewEmptyBlock(id int) Block {
	return &DefaultBlock{Id: id}
}
