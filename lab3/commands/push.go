package commands

import "fmt"

type PushCommand struct {
	code   Code
	length int
}

func (cmd PushCommand) GetOpCode() Code {
	return cmd.code
}

func (cmd PushCommand) Length() int {
	return cmd.length
}

func (cmd PushCommand) String() string {
	return cmd.code.String()
}

type PushByteCommand struct {
	PushCommand
	Arg byte
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{PushCommand: PushCommand{code: PushByte, length: 2}, Arg: arg}
}

func (cmd PushByteCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushIntCommand struct {
	PushCommand
	Arg int32
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{PushCommand: PushCommand{code: PushInt, length: 5}, Arg: arg}
}

func (cmd PushIntCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushLongCommand struct {
	PushCommand
	Arg int64
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{PushCommand: PushCommand{code: PushLong, length: 9}, Arg: arg}
}

func (cmd PushLongCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushReferenceCommand struct {
	PushCommand
	Arg int32
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{PushCommand: PushCommand{code: PushReference, length: 5}, Arg: arg}
}

func (cmd PushReferenceCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}
