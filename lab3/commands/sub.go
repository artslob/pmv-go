package commands

import "strconv"

type SubByteCommand struct {
	ByteCommand
}

func NewSubByteCommand(arg byte) *SubByteCommand {
	return &SubByteCommand{
		ByteCommand: ByteCommand{
			BaseCommand: BaseCommand{
				code:        SubByte,
				argAsString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type SubIntCommand struct {
	Int32Command
}

func NewSubIntCommand(arg int32) *SubIntCommand {
	return &SubIntCommand{
		Int32Command: Int32Command{
			BaseCommand: BaseCommand{
				code:        SubInt,
				argAsString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type SubLongCommand struct {
	Int64Command
}

func NewSubLongCommand(arg int64) *SubLongCommand {
	return &SubLongCommand{
		Int64Command: Int64Command{
			BaseCommand: BaseCommand{
				code:        SubLong,
				argAsString: strconv.FormatInt(arg, 10),
			},
			Arg: arg,
		},
	}
}
