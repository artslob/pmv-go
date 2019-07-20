package commands

type AddByteCommand struct {
	TwoChildCommand
}

func NewAddByteCommand(right, left Command) *AddByteCommand {
	return &AddByteCommand{NewTwoChildCommand(AddByte, left, right)}
}

type AddIntCommand struct {
	TwoChildCommand
}

func NewAddIntCommand(right, left Command) *AddIntCommand {
	return &AddIntCommand{NewTwoChildCommand(AddInt, left, right)}
}

type AddLongCommand struct {
	TwoChildCommand
}

func NewAddLongCommand(right, left Command) *AddLongCommand {
	return &AddLongCommand{NewTwoChildCommand(AddLong, left, right)}
}
