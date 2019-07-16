package blocks

import "fmt"

type Function struct {
	DefaultBlock
}

func NewFunction(id int, text string) *Function {
	return &Function{DefaultBlock: DefaultBlock{Id: id, Text: text}}
}

func (b *Function) String() string {
	return fmt.Sprintf("%2d [label=\"%s\",shape=parallelogram,color=lightblue,style=filled]\n", b.Id, b.Text)
}

func (b *Function) GetAddress() Block {
	return b
}

func (b *Function) SetBranch(Block) {
	panic("function block do not provide branching")
}
