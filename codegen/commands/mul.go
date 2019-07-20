package commands

type MulByteCommand struct {
	TwoChildCommand
}

func NewMulByteCommand(right, left Command) *MulByteCommand {
	return &MulByteCommand{NewTwoChildCommand(MulByte, left, right)}
}

type MulIntCommand struct {
	TwoChildCommand
}

func NewMulIntCommand(right, left Command) *MulIntCommand {
	return &MulIntCommand{NewTwoChildCommand(MulInt, left, right)}
}

type MulLongCommand struct {
	TwoChildCommand
}

func NewMulLongCommand(right, left Command) *MulLongCommand {
	return &MulLongCommand{NewTwoChildCommand(MulLong, left, right)}
}
