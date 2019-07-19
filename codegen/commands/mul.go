package commands

type MulByteCommand struct {
	BaseCommand
}

func NewMulByteCommand() *MulByteCommand {
	return &MulByteCommand{BaseCommand: BaseCommand{code: MulByte}}
}

type MulIntCommand struct {
	BaseCommand
}

func NewMulIntCommand() *MulIntCommand {
	return &MulIntCommand{BaseCommand: BaseCommand{code: MulInt}}
}

type MulLongCommand struct {
	BaseCommand
}

func NewMulLongCommand() *MulLongCommand {
	return &MulLongCommand{BaseCommand: BaseCommand{code: MulLong}}
}
