package parser

import (
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   []string
		tok ItemToken
		lit string
	}{
		{s: []string{``}, tok: EOF, lit: ""},
		{s: []string{``, ``}, tok: EOL, lit: ""},
		{s: []string{`  `}, tok: WS, lit: "  "},
		{s: []string{`.`}, tok: Illegal, lit: "."},
		{s: []string{`@`}, tok: Attr, lit: "@"},
		{s: []string{`Command`}, tok: Command, lit: "Command"},
		{s: []string{`Node`}, tok: Node, lit: "Node"},
		{s: []string{`MetaType`}, tok: MetaType, lit: "MetaType"},
		{s: []string{`Is`}, tok: Is, lit: "Is"},
		{s: []string{`Parent`}, tok: Parent, lit: "Parent"},
		{s: []string{`Param`}, tok: Param, lit: "Param"},
		{s: []string{`Return`}, tok: Return, lit: "Return"},
		{s: []string{`Column`}, tok: Column, lit: "Column"},
		{s: []string{`Value`}, tok: Value, lit: "Value"},
		{s: []string{`Display_Name`}, tok: Ident, lit: "Display_Name"},
	}

	for i, tt := range tests {
		s := NewScanner(tt.s)
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

func TestScanner_ScanText(t *testing.T) {
	var tests = []struct {
		s   []string
		tok ItemToken
		lit string
	}{
		{s: []string{`Sample String`}, tok: Text, lit: "Sample String"},
		{s: []string{`Sample`, `String`}, tok: Text, lit: "Sample String"},
		{s: []string{`Sample`, `String`, `Example`}, tok: Text, lit: "Sample String Example"},
		{s: []string{`Sample`, ``, `String`}, tok: Text, lit: "Sample"},
		{s: []string{`Sample`, `@String`}, tok: Text, lit: "Sample"},
		{s: []string{" \tSample String"}, tok: Text, lit: "Sample String"},
	}

	for i, tt := range tests {
		s := NewScanner(tt.s)
		tok, lit := s.ScanText()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
