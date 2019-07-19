package commands

type PushByteCommand struct {
	ByteCommand
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{ByteCommand: *NewByteCommand(PushByte, arg)}
}

type PushIntCommand struct {
	Int32Command
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{Int32Command: *NewInt32Command(PushInt, arg)}
}

type PushLongCommand struct {
	Int64Command
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{Int64Command: *NewInt64Command(PushLong, arg)}
}

type PushReferenceCommand struct {
	Int32Command
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{Int32Command: *NewInt32Command(PushReference, arg)}
}
