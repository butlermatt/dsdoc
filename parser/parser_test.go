package parser

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	var tests = []struct{
		s []string
		doc *Document
		err string
	}{
		{
			s: []string{
				`@Command Add_Device`,
				`@Is addDevice`,
				`@Parent root`,
				``,
				`Adds a Device to the link`,
				``,
				`This is a long description.`,
				`It really doesn't contain anything special.`,
				`But it is multiline`,
			},
			doc: &Document{
				Type: ActionDoc,
				Name: "Add_Device",
				Is: "addDevice",
				ParentName: "root",
				Short: "Adds a Device to the link",
				Long: "This is a long description. It really doesn't contain anything special. But it is multiline",
			},
		},
	}

	for i, tt := range tests {
		d, err := NewParser(tt.s).Parse()
		var es string
		if err == nil {
			es = ""
		} else {
			es = err.Error()
		}
		if d.Type != tt.doc.Type {
			t.Errorf("%d. Doc Types do not match: exp=%q got=%q", i, tt.doc.Type, d.Type)
		}
		if d.Is != tt.doc.Is {
			t.Errorf("%d. Doc IsType does not match: exp=%q got=%q", i, tt.doc.Is, d.Is)
		}
		if d.ParentName != tt.doc.ParentName {
			t.Errorf("%d. Doc Parent does not match: exp=%q got=%q", i, tt.doc.ParentName, d.ParentName)
		}
		if d.Short != tt.doc.Short {
			t.Errorf("%d. Short Description does not match:\n  exp=%q\n  got=%q\n", i, tt.doc.Short, d.Short)
		}
		if d.Long != tt.doc.Long {
			t.Errorf("%d. Long Description does not match:\n  exp=%q\n  got=%q\n", i, tt.doc.Long, d.Long)
		}
		if es != tt.err {
			t.Errorf("%d. Error mismatch:\n  exp=%q\n  got=%q\n", i, tt.err, es)
		}
	}
}
