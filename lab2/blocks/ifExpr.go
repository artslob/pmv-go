package blocks

type BranchBlock struct {
	DefaultBlock
	branch Block
}

func (b *BranchBlock) GetAddress() Block {
	return b
}

func (b *BranchBlock) SetBranch(block Block) {
	b.branch = block.GetAddress()
}

func (b *BranchBlock) GetBranch() Block {
	return b.branch
}
