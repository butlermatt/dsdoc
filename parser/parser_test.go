package parser

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	var tests = []struct {
		s   []string
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
				``,
				`@Param deviceName string Name of the device to add. It will`,
				`appear as a node on the root of the link.`,
				`@Param username string The Username to access the device.`,
				``,
				`@Return value`,
				`@Column success bool Returns true on success. False otherwise.`,
			},
			doc: &Document{
				Type:       ActionDoc,
				Name:       "Add_Device",
				Is:         "addDevice",
				ParentName: "root",
				Short:      "Adds a Device to the link",
				Long:       "This is a long description. It really doesn't contain anything special. But it is multiline",
				Params: []*Parameter{
					{
						Name:        "deviceName",
						Type:        "string",
						Description: "Name of the device to add. It will appear as a node on the root of the link.",
					},
					{
						Name:        "username",
						Type:        "string",
						Description: "The Username to access the device.",
					},
				},
				Return: "value",
				Columns: []*Parameter{
					{
						Name:        "success",
						Type:        "bool",
						Description: "Returns true on success. False otherwise.",
					},
				},
			},
		},
		{
			s: []string{
				`@Node`,
				`@MetaType test`,
				`@Parent root`,
				``,
				`Short Test node`,
				``,
				`Also has a long description. But no value.`,
			},
			doc: &Document{
				Type:       NodeDoc,
				Name:       "test",
				ParentName: "root",
				Short:      "Short Test node",
				Long:       "Also has a long description. But no value.",
			},
		},
		{
			s: []string{
				`@Node version`,
				`@Parent root`,
				``,
				`Short version description`,
				``,
				`@Value string`,
			},
			doc: &Document{
				Type:       NodeDoc,
				Name:       "version",
				ParentName: "root",
				Short:      "Short version description",
				ValueType:  "string",
			},
		},
	}

	for i, tt := range tests {
		parser := NewParser()
		err := parser.Parse(tt.s)
		if err != nil {
			t.Errorf("%d. Unexpected error %q", i, err)
		}
		doc, err := parser.Build()
		if err != nil {
			t.Errorf("%d. Unexpected error %q", i, err)
		}
		d := doc.Children[0]

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

		if len(tt.doc.Params) != 0 {
			if len(tt.doc.Params) != len(d.Params) {
				t.Errorf("%d. Unequal Parameter count: exp=%d got=%d", i, len(tt.doc.Params), len(d.Params))
			}
			for j, p := range tt.doc.Params {
				tp := d.Params[j]
				if p.Type != tp.Type {
					t.Errorf("%d. Param %d. Param type mismatch: exp=%q got=%q", i, j, p.Type, tp.Type)
				}
				if p.Name != tp.Name {
					t.Errorf("%d. Param %d. Param name mismatch: exp=%q got=%q", i, j, p.Name, tp.Name)
				}
				if p.Description != tp.Description {
					t.Errorf("%d. Param %d. Param description mismatch:\n  exp=%q\n  got=%q", i, j, p.Description, tp.Description)
				}
			}
		} else if len(d.Params) != 0 {
			t.Errorf("%d. Expected 0 parameters, found=%d", i, len(d.Params))
		}

		if d.Return != tt.doc.Return {
			t.Errorf("%d. Return type does not match: exp=%q got=%q", i, tt.doc.Return, d.Return)
		}

		if len(tt.doc.Columns) != 0 {
			if len(tt.doc.Columns) != len(d.Columns) {
				t.Errorf("%d. Unequal Columns count. exp=%d got=%d", i, len(tt.doc.Columns), len(d.Columns))
			}
			for j, p := range tt.doc.Columns {
				tp := d.Columns[i]
				if p.Type != tp.Type {
					t.Errorf("%d. Column %d. Column type mismatch: exp=%q got=%q", i, j, p.Type, tp.Type)
				}
				if p.Name != tp.Name {
					t.Errorf("%d. Column %d. Column name mismatch: exp=%q got=%q", i, j, p.Name, tp.Name)
				}
				if p.Description != tp.Description {
					t.Errorf("%d. Column %d. Column description mismatch:\n  exp=%q\n  got=%q", i, j, p.Description, tp.Description)
				}
			}
		} else if len(d.Columns) != 0 {
			t.Errorf("%d. Expect 0 columns, found=%d", i, len(d.Columns))
		}

		if d.ValueType != tt.doc.ValueType {
			t.Errorf("%d. Value type does not match: exp=%q got=%q", i, tt.doc.ValueType, d.ValueType)
		}

		if es != tt.err {
			t.Errorf("%d. Error mismatch:\n  exp=%q\n  got=%q\n", i, tt.err, es)
		}
	}
}
