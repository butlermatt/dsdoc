package parser

import "unicode/utf8"

const (
	eof = rune(0)
	eol = rune(3)
)

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
		s.line += 1
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
