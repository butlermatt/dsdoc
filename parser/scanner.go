package parser

import "unicode/utf8"

const (
	eof = rune(0)
	eol = rune(3)
)

// AttrChar is the character which denotes a DsDoc Attribute
const AttrChar = "@"

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
