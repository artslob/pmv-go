package types

type Type byte

const (
	Byte Type = iota
	Int
	Long
	Reference
	NewVar
	Bool
	Store
)
