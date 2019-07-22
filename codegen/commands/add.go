package commands

import "github.com/artslob/pmv-go/codegen/types"

type AddByteCommand struct {
	TwoChildCommand
}

func NewAddByteCommand(right, left Command) *AddByteCommand {
	return &AddByteCommand{NewTwoChildCommand(AddByte, types.Byte, left, right)}
}

type AddIntCommand struct {
	TwoChildCommand
}

func NewAddIntCommand(right, left Command) *AddIntCommand {
	return &AddIntCommand{NewTwoChildCommand(AddInt, types.Int, left, right)}
}

type AddLongCommand struct {
	TwoChildCommand
}

func NewAddLongCommand(right, left Command) *AddLongCommand {
	return &AddLongCommand{NewTwoChildCommand(AddLong, types.Long, left, right)}
}
