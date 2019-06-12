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

type SimpleBlock struct {
	Id   int
	Text string
	next Block
}

func (b *SimpleBlock) GetAddress() Block {
	return b
}

func (b *SimpleBlock) SetNext(next Block) {
	b.next = next.GetAddress()
}

func (b *SimpleBlock) GetNext() Block {
	return b.next
}

func (b *SimpleBlock) SetBranch(Block) {
	panic("simple block do not provide branching")
}

func (b *SimpleBlock) GetBranch() Block {
	return nil
}

func (b *SimpleBlock) GetId() int {
	return b.Id
}

func (b *SimpleBlock) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.Id, b.Text)
}

func NewEmptyBlock(id int) Block {
	return &SimpleBlock{Id: id}
}

type IfExpr struct {
	SimpleBlock
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
	SimpleBlock
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
