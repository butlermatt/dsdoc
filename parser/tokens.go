package parser

// Token represents a lexical token.
type ItemToken int

const (
	// Special Tokens
	Illegal ItemToken = iota
	Eof
	Eol
	Ws

	// Literals
	Ident // fields, Table_name
	Text // short/long descriptions.
	ValType // ValueType (number, boolean, value, table, stream, etc)

	// Misc Characters
	Attr // @

	// KEYWORDS
	Command
	Node
	MetaType
	Is
	Parent
	Param
	Return
	Column
	Value
)
