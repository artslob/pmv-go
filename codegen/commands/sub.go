package commands

type SubByteCommand struct {
	TwoChildCommand
}

func NewSubByteCommand(right, left Command) *SubByteCommand {
	return &SubByteCommand{NewTwoChildCommand(SubByte, left, right)}
}

type SubIntCommand struct {
	TwoChildCommand
}

func NewSubIntCommand(right, left Command) *SubIntCommand {
	return &SubIntCommand{NewTwoChildCommand(SubInt, left, right)}
}

type SubLongCommand struct {
	TwoChildCommand
}

func NewSubLongCommand(right, left Command) *SubLongCommand {
	return &SubLongCommand{NewTwoChildCommand(SubLong, left, right)}
}
