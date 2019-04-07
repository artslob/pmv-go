package lab2

import "fmt"

type BlockClass int

const (
	DEFAULT BlockClass = iota
	START
	END
)

type Block struct {
	id    int
	text  string
	next  *Block
	class BlockClass
}

func (b *Block) String() string {
	label := ""
	switch b.class {
	case START:
		label = "START"
	case END:
		label = "END"
	default:
		label = b.text
	}
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.id, label)
}
