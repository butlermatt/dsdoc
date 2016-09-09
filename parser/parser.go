package parser

import (
	"errors"
	"fmt"
)

// DocType represents the type of document being parsed.
type DocType int

const (
	// LinkDoc is a document of the link itself.
	LinkDoc DocType = iota
	// NodeDoc is a document of node type.
	NodeDoc
	// ActionDoc is a document of an action.
	ActionDoc
)

func (d DocType) String() string {
	switch d {
	case LinkDoc:
		return "Link"
	case NodeDoc:
		return "Node"
	case ActionDoc:
		return "Action"
	}
	return ""
}

// Document is the primary container of the DsDoc.
type Document struct {
	Type       DocType
	Path       string
	Name       string
	MetaName   string
	Is         string
	ParentName string
	Parent     *Document
	Children   []*Document
	Short      string
	Long       string
	Params     []*Parameter
	Return     string
	Columns    []*Parameter
	ValueType  string
}

// Parameter is a component of a Action type. Used as either a action
// parameter or return column.
type Parameter struct {
	Name        string
	Type        string
	Description string
}

// Parser represents a parser, which extends the functionality of Scanner
type Parser struct {
	s   *Scanner
	c   map[string]*Document
	r   *Document
	buf struct {
		tok ItemToken
		lit string
		b   bool
	}
}

// NewParser returns a new instance of Parser
func NewParser() *Parser {
	root := &Document{
		Type:  NodeDoc,
		Name:  "root",
		Short: "Root node of the DsLink",
	}
	return &Parser{
		c: map[string]*Document{root.Name: root},
		r: root,
	}
}

// Parse will take the input string slice and try to parse the information.
func (p *Parser) Parse(in []string) error {
	p.s = NewScanner(in)
	p.buf.b = false
	doc := &Document{}

	// First token should be an Attribute character.
	if tok, lit := p.scan(); tok != Attr {
		return fmt.Errorf("found %q, expected %q", lit, AttrChar)
	}

	// Expect DsDoc to start with either @Action, @Node or @Link
	tok, lit := p.scan()
	switch tok {
	case Action:
		doc.Type = ActionDoc
	case Node:
		doc.Type = NodeDoc
	case Link:
		doc.Type = LinkDoc
	default:
		return fmt.Errorf("Expect DocType, found %q", lit)
	}

	if tok, lit = p.scanIgnoreWs(); tok == Ident {
		doc.Path = lit
		doc.Name = lit
		doc.MetaName = lit
	} else if tok == EOF {
		return errors.New("DsDoc unexpectedly terminated early.")
	} else if tok != EOL {
		return fmt.Errorf("Expected ident string or EOL, found %q", lit)
	}

	for {
		p.maybeEol() // Skip any possible end of lines.
		tok, lit = p.scan()

		// Check for EOL again, if so may be start of Short/Long description
		if tok == EOF {
			break
		} else if tok == EOL {
			r := p.s.peak()
			if r != AttrChar {
				tok, lit = p.scanText()
				if tok == Text {
					if doc.Short == "" {
						doc.Short = lit
					} else {
						doc.Long = lit
					}
				}
			}
		} else if tok == Attr {
			tok, lit = p.scan()
			var err error
			switch tok {
			case MetaType:
				err = p.scanMetaType(doc)
			case Is:
				err = p.scanIs(doc)
			case Parent:
				err = p.scanParent(doc)
			case Param:
				err = p.scanParam(doc)
			case Return:
				err = p.scanReturn(doc)
			case Column:
				err = p.scanColumn(doc)
			case Value:
				err = p.scanValue(doc)
			}

			if err != nil {
				return err
			}
		}
	}

	// TODO: Verify required values are set
	if doc.Name == "" {
		return errors.New("DsDoc missing required Name or MetaType field")
	}
	ed := p.c[doc.MetaName]
	if ed != nil {
		return fmt.Errorf("DsDoc with meta name %q already exists", doc.MetaName)
	}

	if doc.ParentName == "" {
		return errors.New("DsDoc missing required Parent field")
	}

	pd := p.c[doc.ParentName]
	if pd != nil {
		doc.Parent = pd
		pd.Children = append(pd.Children, doc)
	}

	p.c[doc.MetaName] = doc
	return nil
}

// Build completes the final linking of documents and returns the root document.
func (p *Parser) Build() (*Document, error) {
	for key, doc := range p.c {
		if key == "root" {
			continue
		}
		if doc.Parent == nil {
			pd, ok := p.c[doc.ParentName]
			if !ok {
				return nil, fmt.Errorf("Unable to locate Parent named %q referenced by %q", doc.ParentName, key)
			}
			doc.Parent = pd
			pd.Children = append(pd.Children, doc)
		}
	}
	return p.r, nil
}

func (p *Parser) scan() (ItemToken, string) {
	if p.buf.b {
		p.buf.b = false
		return p.buf.tok, p.buf.lit
	}

	tok, lit := p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

func (p *Parser) unscan() { p.buf.b = true }

func (p *Parser) scanIgnoreWs() (ItemToken, string) {
	tok, lit := p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return tok, lit
}

func (p *Parser) scanTypeIgnoreWs() (ItemToken, string) {
	if p.buf.b && p.buf.tok == TypeIdent {
		p.buf.b = false
		return p.buf.tok, p.buf.lit
	}

	tok, lit := p.s.scanTypeIdent()
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

func (p *Parser) scanText() (ItemToken, string) {
	if p.buf.b && p.buf.tok == Text {
		p.buf.b = false
		return p.buf.tok, p.buf.lit
	}

	tok, lit := p.s.ScanText()
	p.buf.tok, p.buf.lit = tok, lit
	return tok, lit
}

func (p *Parser) scanIs(d *Document) error {
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	d.Is = lit

	return nil
}

func (p *Parser) scanMetaType(d *Document) error {
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	if d.Name == "" {
		d.Name = lit
	}
	d.MetaName = lit
	return nil
}

func (p *Parser) scanParent(d *Document) error {
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	d.ParentName = lit
	return nil
}

func (p *Parser) scanParam(d *Document) error {
	param := &Parameter{}
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	param.Name = lit

	tok, lit = p.scanTypeIgnoreWs()
	if tok != TypeIdent {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	param.Type = lit // TODO: Check types in the future.

	tok, lit = p.scanText()
	if tok != Text {
		return fmt.Errorf("Expected Text, found %q (%q)", lit, tok)
	}
	param.Description = lit
	d.Params = append(d.Params, param)

	return nil
}

func (p *Parser) scanReturn(d *Document) error {
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	d.Return = lit
	return nil
}

func (p *Parser) scanColumn(d *Document) error {
	param := &Parameter{}
	tok, lit := p.scanIgnoreWs()
	if tok != Ident {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	param.Name = lit

	tok, lit = p.scanTypeIgnoreWs()
	if tok != TypeIdent {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	param.Type = lit // TODO: Check types in the future.

	tok, lit = p.scanText()
	if tok != Text {
		return fmt.Errorf("Expected Text, found %q (%q)", lit, tok)
	}
	param.Description = lit
	d.Columns = append(d.Columns, param)

	return nil
}

func (p *Parser) scanValue(d *Document) error {
	tok, lit := p.scanTypeIgnoreWs()
	if tok != TypeIdent {
		return fmt.Errorf("Expected Ident, found %q (%q)", lit, tok)
	}
	d.ValueType = lit
	return nil
}

func (p *Parser) maybeEol() {
	if tok, _ := p.scan(); tok != EOL {
		p.unscan()
	}
}
