package blocks

import "fmt"

type Block interface {
	fmt.Stringer
	GetAddress() Block
	GetId() int
	// SetNext should call GetAddress() on his argument
	SetNext(Block)
	GetNext() Block
	// SetBranch should call GetAddress() on his argument
	SetBranch(Block)
	GetBranch() Block
}

type DefaultBlock struct {
	Id   int
	Text string
	next Block
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

type IfExpr struct {
	DefaultBlock
	branch Block
}

func (b *IfExpr) GetAddress() Block {
	return b
}

func (b *IfExpr) SetBranch(block Block) {
	b.branch = block.GetAddress()
}

func (b *IfExpr) GetBranch() Block {
	return b.branch
}

type IfBlock struct {
	DefaultBlock
	Expr, Then, Else_, End Block
}

func (b *IfBlock) GetAddress() Block {
	return b.Expr.GetAddress()
}

func (b *IfBlock) SetNext(block Block) {
	b.End.SetNext(block.GetAddress())
}

func (b *IfBlock) GetNext() Block {
	return b.End.GetNext()
}

func (b *IfBlock) SetBranch(Block) {
	panic("if block do not provide branching")
}
