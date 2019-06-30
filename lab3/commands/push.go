package commands

import "fmt"

type PushByteCommand struct {
	Arg byte
}

func (cmd PushByteCommand) GetOpCode() Code {
	return PushByte
}

func (cmd PushByteCommand) Length() int {
	return 2
}

func (cmd PushByteCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushIntCommand struct {
	Arg int32
}

func (cmd PushIntCommand) GetOpCode() Code {
	return PushInt
}

func (cmd PushIntCommand) Length() int {
	return 5
}

func (cmd PushIntCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushLongCommand struct {
	Arg int64
}

func (cmd PushLongCommand) GetOpCode() Code {
	return PushLong
}

func (cmd PushLongCommand) Length() int {
	return 9
}

func (cmd PushLongCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}

type PushReferenceCommand struct {
	PushIntCommand
}

func (cmd PushReferenceCommand) GetOpCode() Code {
	return PushReference
}

func (cmd PushReferenceCommand) String() string {
	return fmt.Sprintf("%s %d", cmd.GetOpCode().String(), cmd.Arg)
}
