package parser

import "fmt"

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
	// TypeIdent indicates a type specific identifier token and is less restrictive
	TypeIdent
	// Text indicates a text string such as in short/long description
	Text
	// Attr indicates an DsDoc Attribute character
	Attr // @
	// Action is a DsDoc attribute keyword.
	Action
	// Node is a DsDoc attribute keyword.
	Node
	// Link is a DsDoc attribute keyword.
	Link
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

func (i ItemToken) String() string {
	t := "UNKNOWN"
	switch i {
	case Illegal:
		t = "Illegal"
	case EOF:
		t = "EOF"
	case EOL:
		t = "EOL"
	case WS:
		t = "Whitespace"
	case Ident:
		t = "Ident"
	case TypeIdent:
		t = "TypeIdent"
	case Text:
		t = "Text"
	case Attr:
		t = string(AttrChar)
	}
	return fmt.Sprintf("Token(%v)", t)
}
