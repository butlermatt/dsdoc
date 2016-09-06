package docs

import (
	"fmt"
	"strings"
)

// Prefix is the DsDoc comment style.
const Prefix string = "//*"

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

// Parameter is a component of a command type. Used as either a command
// parameter or return column.
type Parameter struct {
	Name        string
	Type        string
	Description string
}

// Document is the primary container of the DsDocument.
type Document struct {
	Type       DocType
	Name       string
	Is         string
	ParentName string
	parent     *Document
	children   []*Document
	Short      string
	Long       string
	Params     []*Parameter
	Return     string
	Columns    []*Parameter
}

// TrimPrefix takes the slice of strings which contain the DsDoc comment.
// It will skip any lines which do not have
func TrimPrefix(s []string) ([]string, error) {
	r := []string{}
	for _, str := range s {
		if !strings.HasPrefix(str, Prefix) {
			return nil, fmt.Errorf("Invalid DsDoc. Cannot contain a line without a DsDoc comment")
		}
		tmp := strings.TrimSpace(strings.Trim(str, Prefix))
		r = append(r, tmp)
	}

	return r, nil
}

// New creates a new Document from a slice of strings which compose the DsDoc
func New(blob []string) (*Document, error) {
	d := new(Document)
	// TODO: Fill out the body
	lines, err := TrimPrefix(blob)
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("DsDoc contains no data")
	}

	typeStr := strings.Fields(lines[0])
	switch {
	case strings.EqualFold(typeStr[0], "@Command"):
		d.Type = ActionDoc
	case strings.EqualFold(typeStr[0], "@Node"):
		d.Type = NodeDoc
	case strings.EqualFold(typeStr[0], "@Link"):
		d.Type = LinkDoc
	default:
		return nil, fmt.Errorf("DsDoc is missing the type declaration")
	}

	if len(typeStr) >= 2 {
		d.Name = typeStr[1]
	}

	return d, nil
}
