package blocks

import "fmt"

type BranchBlock struct {
	DefaultBlock
	branch Block
}

func (b *BranchBlock) String() string {
	return fmt.Sprintf("%2d [label=\"%s\",shape=diamond]\n", b.Id, b.Text)
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
