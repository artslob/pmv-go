package operations

type Code byte

const (
	Nop  Code = iota
	Halt      = iota

	Neg = iota
	Not = iota

	Mul = iota
	Div = iota
	Mod = iota

	Add = iota
	Sub = iota

	IfEqZero     = iota
	IfNotEqZero  = iota
	IfLessZero   = iota
	IfLessEqZero = iota
	IfGtZero     = iota
	IfGtEqZero   = iota

	IfEq     = iota
	IfNotEq  = iota
	IfLess   = iota
	IfLessEq = iota
	IfGt     = iota
	IfGtEq   = iota

	And = iota
	Or  = iota
	Xor = iota

	PushChar = iota
	PushInt  = iota

	StoreChar = iota
	StoreInt  = iota
)
