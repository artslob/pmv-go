package blocks

type IfBlock struct {
	DefaultBlock
	Expr, Then, Else_, End Block
}

func (b *IfBlock) String() string {
	panic("if block does not support Stringer")
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
