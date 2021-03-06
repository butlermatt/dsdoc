package main

import (
	"bytes"
	"fmt"
	"github.com/butlermatt/dsdoc/parser"
	"strings"
	"sort"
)

var buf bytes.Buffer
var tree bytes.Buffer

type ByAction []*parser.Document
func (a ByAction) Len() int { return len(a) }
func (a ByAction) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAction) Less(i, j int) bool { return a[i].Type == parser.ActionDoc && a[j].Type != parser.ActionDoc }

func genText(doc *parser.Document) bytes.Buffer {
	walkTextDoc(doc, "")
	tree.WriteString("\n---\n\n")
	tree.WriteString(buf.String())
	return tree
}

func walkTextDoc(doc *parser.Document, sep string) {
	buf.WriteString(fmt.Sprintln("Name:", doc.Name))
	buf.WriteString(fmt.Sprint("\n", doc.Short, "\n\n"))
	buf.WriteString(fmt.Sprintln("Type:", doc.Type))
	if doc.Is != "" {
		buf.WriteString(fmt.Sprintln("$is:", doc.Is))
	}
	if doc.ParentName != "" {
		buf.WriteString(fmt.Sprintln("Parent:", doc.Parent.Name))
	}
	if doc.Long != "" {
		buf.WriteString(fmt.Sprint("Description:\n", doc.Long, "\n\n"))
	}

	if doc.Type == parser.ActionDoc {
		if len(doc.Params) > 0 {
			buf.WriteString("Params:\n")
			for _, p := range doc.Params {
				buf.WriteString(fmt.Sprintln("     Name:", p.Name))
				buf.WriteString(fmt.Sprintln("     Type:", p.Type))
				buf.WriteString(fmt.Sprintln("    ", p.Description))
				buf.WriteRune('\n')
			}
			buf.WriteRune('\n')
		}

		buf.WriteString(fmt.Sprintln("Return type:", doc.Return))
		if len(doc.Columns) > 0 {
			buf.WriteString("Columns:\n")
			for _, p := range doc.Columns {
				buf.WriteString(fmt.Sprintln("     Name:", p.Name))
				buf.WriteString(fmt.Sprintln("     Type:", p.Type))
				buf.WriteString(fmt.Sprintln("    ", p.Description))
				buf.WriteRune('\n')
			}
		}
	}

	if doc.ValueType != "" {
		buf.WriteString(fmt.Sprintln("Value Type:", doc.ValueType, "  "))
		buf.WriteString(fmt.Sprintln("Writable:", doc.Writable, "  "))
	}
	buf.WriteString("\n---\n\n")

	if (doc.Type == parser.ActionDoc) {
		var args string
		var params []string
		for _, a := range doc.Params {
			params = append(params, a.Name)
		}
		args = strings.Join(params, ", ")
		tree.WriteString(fmt.Sprintf("%s- @%s(%s)\n", sep, doc.Name, args))
	} else {
		var vType string
		if doc.ValueType != "" {
			vType = fmt.Sprintf(" *%s (%s)*", doc.ValueType, doc.Writable)
		}
		tree.WriteString(fmt.Sprintf("%s- %s%s\n", sep, doc.Name, vType))
	}
	if len(doc.Children) > 0 {
		sort.Stable(ByAction(doc.Children))
		for _, ch := range doc.Children {
			walkTextDoc(ch, sep+" |")
		}
	}
}

func genMarkdown(doc *parser.Document) bytes.Buffer {
	tree.WriteString(" <pre>\n")
	walkMdDoc(doc, "")
	tree.WriteString(" </pre>\n\n---\n\n")
	tree.WriteString(buf.String())
	return tree
}

func walkMdDoc(doc *parser.Document, sep string) {
	buf.WriteString(fmt.Sprint("### ", doc.Name, "  \n\n"))
	buf.WriteString(fmt.Sprint(doc.Short, "  \n\n"))
	buf.WriteString(fmt.Sprintln("Type:", doc.Type, "  "))
	if doc.Is != "" {
		buf.WriteString(fmt.Sprintln("$is:", doc.Is, "  "))
	}
	if doc.ParentName != "" {
		buf.WriteString(fmt.Sprintf("Parent: [%s](#%s)  \n", doc.Parent.Name, strings.ToLower(doc.Parent.Name)))
	}
	if doc.Long != "" {
		buf.WriteString(fmt.Sprint("\nDescription:  \n", doc.Long, "  \n\n"))
	}

	if doc.Type == parser.ActionDoc {
		if len(doc.Params) > 0 {
			buf.WriteString("Params:  \n\n")
			buf.WriteString("Name | Type | Description\n")
			buf.WriteString("--- | --- | ---\n")
			for _, p := range doc.Params {
				buf.WriteString(fmt.Sprintf("%s | `%s` | %s\n", p.Name, p.Type, p.Description))
			}
			buf.WriteString("\n")
		}

		buf.WriteString(fmt.Sprintln("Return type:", doc.Return, "  "))
		if len(doc.Columns) > 0 {
			buf.WriteString("Columns:  \n\n")
			buf.WriteString("Name | Type | Description\n")
			buf.WriteString("--- | --- | ---\n")
			for _, p := range doc.Columns {
				buf.WriteString(fmt.Sprintf("%s | `%s` | %s \n", p.Name, p.Type, p.Description))
			}
		}
	}

	if doc.ValueType != "" {
		buf.WriteString(fmt.Sprintf("Value Type: `%s`  \n", doc.ValueType))
		buf.WriteString(fmt.Sprintf("Writable: `%s`  \n", doc.Writable))
	}
	buf.WriteString("\n---\n\n")

	if (doc.Type == parser.ActionDoc) {
		var args string
		var params []string
		for _, a := range doc.Params {
			params = append(params, a.Name)
		}
		args = strings.Join(params, ", ")
		tree.WriteString(fmt.Sprintf("%s-[@%s(%s)](#%s)\n", sep, doc.Name, args, strings.ToLower(doc.Name)))
	} else {
		var vType string
		if doc.ValueType != "" {
			vType = fmt.Sprintf(" - %s", doc.ValueType)
		}
		tree.WriteString(fmt.Sprintf("%s-[%s](#%s)%s\n", sep, doc.Name, strings.ToLower(doc.Name), vType))
	}
	if len(doc.Children) > 0 {
		sort.Stable(ByAction(doc.Children))
		for _, ch := range doc.Children {
			walkMdDoc(ch, sep+" |")
		}
	}
}
