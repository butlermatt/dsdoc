package parser

import (
	"bytes"
	"unicode/utf8"
)

const (
	eof = rune(0)
	eol = rune(3)
)

// AttrChar is the character which denotes a DsDoc Attribute
const AttrChar = '@'

// Scanner represents a lexical scanner.
type Scanner struct {
	in    []string
	line  int
	start int
	pos   int
	width int
}

// read will return the next rune in the input.
func (s *Scanner) read() rune {
	if s.line >= len(s.in) {
		return eof
	}
	if s.pos >= len(s.in[s.line]) {
		s.line++
		s.pos = 0
		s.width = 0
		if s.line >= len(s.in) {
			return eof
		}
		return eol
	}
	var r rune
	r, s.width = utf8.DecodeRuneInString(s.in[s.line][s.pos:])
	s.pos += s.width
	return r
}

// unread places the previous rune back in the reader.
func (s *Scanner) unread() {
	if s.pos == 0 {
		s.line--
		s.pos = len(s.in[s.line])
	} else {
		s.pos -= s.width
	}
}

// Scan returns the next token and the literal value.
func (s *Scanner) Scan() (ItemToken, string) {
	r := s.read()

	if isWs(r) {
		s.unread()
		return s.scanWhitespace()
	} else if isAlphaNum(r) {
		s.unread()
		return s.scanIdent()
	}

	switch r {
	case eof:
		return EOF, ""
	case eol:
		return EOL, ""
	case AttrChar:
		return Attr, string(r)
	}
	return Illegal, string(r)
}

// ScanText returns a Text token and consumes all text until two sequential eol runes are reached.
func (s *Scanner) ScanText() (ItemToken, string) {
	var buf bytes.Buffer

	// Trim leading whitespace
	r := s.read();
	for ; isWs(r); r = s.read() {
	}

	buf.WriteRune(r)

	for {
		if r := s.read(); r == eof {
			s.unread()
			break
		} else if r == eol {
			n := s.peak()
			if n == eof || n == eol || n == AttrChar {
				s.unread()
				break
			}
			buf.WriteRune(' ') //
		} else {
			buf.WriteRune(r)
		}
	}

	return Text, buf.String()
}

// scanWhitespace consumes all contiguous whitespace.
func (s *Scanner) scanWhitespace() (ItemToken, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if r := s.read(); r == eof || r == eol {
			s.unread()
			break
		} else if !isWs(r) {
			s.unread()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes all contiguous ident runes.
func (s *Scanner) scanIdent() (ItemToken, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if r := s.read(); r == eof || r == eol {
			s.unread()
			break
		} else if !isAlphaNum(r) && r != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	switch buf.String() {
	case "Command":
		return Command, buf.String()
	case "Node":
		return Node, buf.String()
	case "Link":
		return Link, buf.String()
	case "MetaType":
		return MetaType, buf.String()
	case "Is":
		return Is, buf.String()
	case "Parent":
		return Parent, buf.String()
	case "Param":
		return Param, buf.String()
	case "Return":
		return Return, buf.String()
	case "Column":
		return Column, buf.String()
	case "Value":
		return Value, buf.String()
	}

	return Ident, buf.String()
}

// peak returns, but does not consume, the next rune.
func (s *Scanner) peak() rune {
	r := s.read()
	s.unread()
	return r
}

// NewScanner returns a new instance of Scanner.
func NewScanner(in []string) *Scanner {
	return &Scanner{in: in}
}

func isWs(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlphaNum(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
