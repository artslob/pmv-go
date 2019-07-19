package commands

type SubByteCommand struct {
	BaseCommand
}

func NewSubByteCommand() *SubByteCommand {
	return &SubByteCommand{BaseCommand: BaseCommand{code: SubByte}}
}

type SubIntCommand struct {
	BaseCommand
}

func NewSubIntCommand() *SubIntCommand {
	return &SubIntCommand{BaseCommand: BaseCommand{code: SubInt}}
}

type SubLongCommand struct {
	BaseCommand
}

func NewSubLongCommand() *SubLongCommand {
	return &SubLongCommand{BaseCommand: BaseCommand{code: SubLong}}
}
