package commands

type AddByteCommand struct {
	BaseCommand
}

func NewAddByteCommand(arg byte) *AddByteCommand {
	return &AddByteCommand{BaseCommand: BaseCommand{code: AddByte}}
}

type AddIntCommand struct {
	BaseCommand
}

func NewAddIntCommand(arg int32) *AddIntCommand {
	return &AddIntCommand{BaseCommand: BaseCommand{code: AddInt}}
}

type AddLongCommand struct {
	BaseCommand
}

func NewAddLongCommand(arg int64) *AddLongCommand {
	return &AddLongCommand{BaseCommand: BaseCommand{code: AddLong}}
}
