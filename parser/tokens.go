package parser

// ItemToken represents a lexical token.
type ItemToken int

const (
	// Illegal indicates an illegal syntax
	Illegal ItemToken = iota
	// EOF indicates end of input
	EOF
	// EOL indicates end of line token
	EOL
	// WS indicates a whitespace token
	WS
	// Ident indicates an identifier token, such as fields, Path_name
	Ident
	// Text indicates a text string such as in short/long description
	Text
	// ValType indicates what the Value's type is.
	ValType
	// Attr indicates an attribute character token
	Attr // @
	// Command is a DsDoc attribute keyword.
	Command
	// Node is a DsDoc attribute keyword.
	Node
	// MetaType is a DsDoc attribute keyword.
	MetaType
	// Is is a DsDoc attribute keyword.
	Is
	// Parent is a DsDoc attribute keyword.
	Parent
	// Param is a DsDoc attribute keyword.
	Param
	// Return is a DsDoc attribute keyword.
	Return
	// Column is a DsDoc attribute keyword.
	Column
	// Value is a DsDoc attribute keyword.
	Value
)
