package commands

import "github.com/artslob/pmv-go/codegen/types"

type PushByteCommand struct {
	ByteCommand
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{ByteCommand: *NewByteCommand(PushByte, types.Byte, arg)}
}

type PushIntCommand struct {
	Int32Command
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{Int32Command: *NewInt32Command(PushInt, types.Int, arg)}
}

type PushLongCommand struct {
	Int64Command
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{Int64Command: *NewInt64Command(PushLong, types.Long, arg)}
}

type PushReferenceCommand struct {
	Int32Command
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{Int32Command: *NewInt32Command(PushReference, types.Reference, arg)}
}
