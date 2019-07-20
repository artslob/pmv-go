package commands

type ModByteCommand struct {
	TwoChildCommand
}

func NewModByteCommand(right, left Command) *ModByteCommand {
	return &ModByteCommand{NewTwoChildCommand(ModByte, left, right)}
}

type ModIntCommand struct {
	TwoChildCommand
}

func NewModIntCommand(right, left Command) *ModIntCommand {
	return &ModIntCommand{NewTwoChildCommand(ModInt, left, right)}
}

type ModLongCommand struct {
	TwoChildCommand
}

func NewModLongCommand(right, left Command) *ModLongCommand {
	return &ModLongCommand{NewTwoChildCommand(ModLong, left, right)}
}
