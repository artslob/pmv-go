package blocks

import "fmt"

type Block interface {
	fmt.Stringer
	GetAddress() Block
	GetId() int
	// SetNext must call GetAddress() on his argument
	SetNext(Block)
	GetNext() Block
	// SetBranch must call GetAddress() on his argument
	SetBranch(Block)
	GetBranch() Block
}
