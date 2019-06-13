package blocks

// Create Group block if number of elements > 0. Otherwise create EmptyBlock.
type Group struct {
	Id         int
	statements []Block
}

func (g *Group) String() string {
	panic("group block does not support Stringer")
}

func (g *Group) GetAddress() Block {
	return g.statements[0].GetAddress()
}

func (g *Group) GetId() int {
	return g.Id
}

func (g *Group) SetNext(block Block) {
	g.statements[len(g.statements)-1].SetNext(block.GetAddress())
}

func (g *Group) GetNext() Block {
	return g.statements[len(g.statements)-1].GetNext()
}

func (g *Group) SetBranch(Block) {
	panic("group block does not support branching")
}

func (g *Group) GetBranch() Block {
	return nil
}

func (g *Group) AddBlock(block Block) {
	// TODO fix order
	if len(g.statements) > 0 {
		g.SetNext(block)
	}
	g.statements = append(g.statements, block)
}
