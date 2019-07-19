package commands

type DivByteCommand struct {
	BaseCommand
}

func NewDivByteCommand() *DivByteCommand {
	return &DivByteCommand{BaseCommand: BaseCommand{code: DivByte}}
}

type DivIntCommand struct {
	BaseCommand
}

func NewDivIntCommand() *DivIntCommand {
	return &DivIntCommand{BaseCommand: BaseCommand{code: DivInt}}
}

type DivLongCommand struct {
	BaseCommand
}

func NewDivLongCommand() *DivLongCommand {
	return &DivLongCommand{BaseCommand: BaseCommand{code: DivLong}}
}
