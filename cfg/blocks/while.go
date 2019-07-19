package blocks

type While struct {
	DefaultBlock
	Expression Block
	Body       Block
	Breaks     []Block
}

func (b *While) String() string {
	panic("While block does not support Stringer")
}

func (b *While) GetAddress() Block {
	return b.Expression.GetAddress()
}

func (b *While) SetNext(block Block) {
	b.Expression.SetNext(block)
	if len(b.Breaks) == 0 {
		return
	}
	for _, br := range b.Breaks {
		br.SetNext(block)
	}
}

func (b *While) GetNext() Block {
	return b.Expression.GetNext()
}

func (b *While) SetBranch(Block) {
	panic("While block do not provide branching")
}
