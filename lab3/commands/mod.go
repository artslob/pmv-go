package commands

type ModByteCommand struct {
	BaseCommand
}

func NewModByteCommand() *ModByteCommand {
	return &ModByteCommand{BaseCommand: BaseCommand{code: ModByte}}
}

type ModIntCommand struct {
	BaseCommand
}

func NewModIntCommand() *ModIntCommand {
	return &ModIntCommand{BaseCommand: BaseCommand{code: ModInt}}
}

type ModLongCommand struct {
	BaseCommand
}

func NewModLongCommand() *ModLongCommand {
	return &ModLongCommand{BaseCommand: BaseCommand{code: ModLong}}
}
