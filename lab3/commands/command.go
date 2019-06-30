package commands

import "fmt"

type Command interface {
	GetOpCode() Code
	Length() int
	fmt.Stringer
}
