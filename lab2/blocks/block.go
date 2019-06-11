package blocks

import "fmt"

type Block interface {
	fmt.Stringer
	GetAddress() Block
	SetNext(Block)
	GetNext() Block
	SetBranch(Block)
	GetBranch() Block
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

func (b *SimpleBlock) SetNext(next Block) {
	b.Next = next
}

func (b *SimpleBlock) GetNext() Block {
	return b.Next
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

type IfExpr struct {
	Id     int
	Text   string
	Next   Block
	Branch Block
	Level  int
}

func (b *IfExpr) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.Id, b.Text)
}

func (b *IfExpr) GetAddress() Block {
	return b
}

func (b *IfExpr) SetNext(block Block) {
	b.Next = block
}

func (b *IfExpr) GetNext() Block {
	return b.Next
}

func (b *IfExpr) SetBranch(block Block) {
	b.Branch = block
}

func (b *IfExpr) GetBranch() Block {
	return b.Branch
}

func (b *IfExpr) GetId() int {
	return b.Id
}

type IfBlock struct {
	Id                     int
	Text                   string
	Next                   Block
	Branch                 Block
	Level                  int
	Expr, Then, Else_, End Block
}

func (b *IfBlock) String() string {
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.Id, b.Text)
}

func (b *IfBlock) GetAddress() Block {
	return b.Expr.GetAddress()
}

func (b *IfBlock) SetNext(block Block) {
	b.End.SetNext(block)
}

func (b *IfBlock) GetNext() Block {
	return b.Next
}

func (b *IfBlock) SetBranch(Block) {
	panic("if block do not provide branching")
}

func (b *IfBlock) GetBranch() Block {
	return nil
}

func (b *IfBlock) GetId() int {
	return b.Id
}
