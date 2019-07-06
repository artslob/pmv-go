package commands

type SubByteCommand struct {
	BaseCommand
}

func NewSubByteCommand(arg byte) *SubByteCommand {
	return &SubByteCommand{BaseCommand: BaseCommand{code: SubByte}}
}

type SubIntCommand struct {
	BaseCommand
}

func NewSubIntCommand(arg int32) *SubIntCommand {
	return &SubIntCommand{BaseCommand: BaseCommand{code: SubInt}}
}

type SubLongCommand struct {
	BaseCommand
}

func NewSubLongCommand(arg int64) *SubLongCommand {
	return &SubLongCommand{BaseCommand: BaseCommand{code: SubLong}}
}
