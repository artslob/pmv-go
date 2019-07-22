package commands

import "github.com/artslob/pmv-go/codegen/types"

type MulByteCommand struct {
	TwoChildCommand
}

func NewMulByteCommand(right, left Command) *MulByteCommand {
	return &MulByteCommand{NewTwoChildCommand(MulByte, types.Byte, left, right)}
}

type MulIntCommand struct {
	TwoChildCommand
}

func NewMulIntCommand(right, left Command) *MulIntCommand {
	return &MulIntCommand{NewTwoChildCommand(MulInt, types.Int, left, right)}
}

type MulLongCommand struct {
	TwoChildCommand
}

func NewMulLongCommand(right, left Command) *MulLongCommand {
	return &MulLongCommand{NewTwoChildCommand(MulLong, types.Long, left, right)}
}
