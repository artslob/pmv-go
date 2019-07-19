package blocks

type Until struct {
	DefaultBlock
	Statement  Block
	Expression Block
	Breaks     []Block
}

func (b *Until) String() string {
	panic("until block does not support Stringer")
}

func (b *Until) GetAddress() Block {
	return b.Statement.GetAddress()
}

func (b *Until) SetNext(block Block) {
	b.Expression.SetNext(block)
	if len(b.Breaks) == 0 {
		return
	}
	for _, br := range b.Breaks {
		br.SetNext(block)
	}
}

func (b *Until) GetNext() Block {
	return b.Expression.GetNext()
}

func (b *Until) SetBranch(Block) {
	panic("until block do not provide branching")
}
