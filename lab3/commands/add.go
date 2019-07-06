package commands

type AddByteCommand struct {
	BaseCommand
}

func NewAddByteCommand() *AddByteCommand {
	return &AddByteCommand{BaseCommand: BaseCommand{code: AddByte}}
}

type AddIntCommand struct {
	BaseCommand
}

func NewAddIntCommand() *AddIntCommand {
	return &AddIntCommand{BaseCommand: BaseCommand{code: AddInt}}
}

type AddLongCommand struct {
	BaseCommand
}

func NewAddLongCommand() *AddLongCommand {
	return &AddLongCommand{BaseCommand: BaseCommand{code: AddLong}}
}
