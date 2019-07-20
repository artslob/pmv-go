package commands

type DivByteCommand struct {
	TwoChildCommand
}

func NewDivByteCommand(right, left Command) *DivByteCommand {
	return &DivByteCommand{NewTwoChildCommand(DivByte, left, right)}
}

type DivIntCommand struct {
	TwoChildCommand
}

func NewDivIntCommand(right, left Command) *DivIntCommand {
	return &DivIntCommand{NewTwoChildCommand(DivInt, left, right)}
}

type DivLongCommand struct {
	TwoChildCommand
}

func NewDivLongCommand(right, left Command) *DivLongCommand {
	return &DivLongCommand{NewTwoChildCommand(DivLong, left, right)}
}
