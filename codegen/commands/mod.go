package commands

import "github.com/artslob/pmv-go/codegen/types"

type ModByteCommand struct {
	TwoChildCommand
}

func NewModByteCommand(right, left Command) *ModByteCommand {
	return &ModByteCommand{NewTwoChildCommand(ModByte, types.Byte, left, right)}
}

type ModIntCommand struct {
	TwoChildCommand
}

func NewModIntCommand(right, left Command) *ModIntCommand {
	return &ModIntCommand{NewTwoChildCommand(ModInt, types.Int, left, right)}
}

type ModLongCommand struct {
	TwoChildCommand
}

func NewModLongCommand(right, left Command) *ModLongCommand {
	return &ModLongCommand{NewTwoChildCommand(ModLong, types.Long, left, right)}
}
