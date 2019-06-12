package blocks

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
