package blocks

// Create Group block if number of elements > 0. Otherwise create EmptyBlock.
//
// Assumes that elements go in reverse order (e.g. from stack)
type Group struct {
	Id         int
	statements []Block
}

func (g *Group) String() string {
	panic("group block does not support Stringer")
}

func (g *Group) GetAddress() Block {
	return g.statements[len(g.statements)-1].GetAddress()
}

func (g *Group) GetId() int {
	return g.Id
}

func (g *Group) SetNext(block Block) {
	g.statements[0].SetNext(block.GetAddress())
}

func (g *Group) GetNext() Block {
	return g.statements[0].GetNext()
}

func (g *Group) SetBranch(Block) {
	panic("group block does not support branching")
}

func (g *Group) GetBranch() Block {
	return nil
}

// Assumes that elements go in reverse order (e.g. from stack)
func (g *Group) AddBlock(block Block) {
	if len(g.statements) > 0 {
		block.SetNext(g.statements[len(g.statements)-1])
	}
	g.statements = append(g.statements, block)
}
