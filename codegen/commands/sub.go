package commands

import "github.com/artslob/pmv-go/codegen/types"

type SubByteCommand struct {
	TwoChildCommand
}

func NewSubByteCommand(right, left Command) *SubByteCommand {
	return &SubByteCommand{NewTwoChildCommand(SubByte, types.Byte, left, right)}
}

type SubIntCommand struct {
	TwoChildCommand
}

func NewSubIntCommand(right, left Command) *SubIntCommand {
	return &SubIntCommand{NewTwoChildCommand(SubInt, types.Int, left, right)}
}

type SubLongCommand struct {
	TwoChildCommand
}

func NewSubLongCommand(right, left Command) *SubLongCommand {
	return &SubLongCommand{NewTwoChildCommand(SubLong, types.Long, left, right)}
}
