package commands

import "github.com/artslob/pmv-go/codegen/types"

type DivByteCommand struct {
	TwoChildCommand
}

func NewDivByteCommand(right, left Command) *DivByteCommand {
	return &DivByteCommand{NewTwoChildCommand(DivByte, types.Byte, left, right)}
}

type DivIntCommand struct {
	TwoChildCommand
}

func NewDivIntCommand(right, left Command) *DivIntCommand {
	return &DivIntCommand{NewTwoChildCommand(DivInt, types.Int, left, right)}
}

type DivLongCommand struct {
	TwoChildCommand
}

func NewDivLongCommand(right, left Command) *DivLongCommand {
	return &DivLongCommand{NewTwoChildCommand(DivLong, types.Long, left, right)}
}
