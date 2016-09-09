package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/butlermatt/dsdoc/parser"
	"github.com/butlermatt/dsdoc/trim"
)

const (
	md = "md"   // Markdown
	tx = "text" //text
)

var ValidFiles = [...]string{
	".dart",
	".java",
	".go",
	".c",
	".cpp",
	".cc",
	".js",
	".ts",
	".es",
}

var psr *parser.Parser

func isSrcFile(ext string) bool {
	for i := 0; i < len(ValidFiles); i++ {
		if ext == ValidFiles[i] {
			return true
		}
	}
	return false
}

func main() {
	var (
		ty = flag.String("t", "md", "output type [md|text]")
		fn = flag.String("o", "api.md", "output file name")
	)

	flag.Parse()
	if *ty != md && *ty != tx {
		fmt.Fprintf(os.Stderr, "Unknown output type: %q\n", *ty)
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	psr = parser.NewParser()

	filepath.Walk(wd, walkFn)

	doc, err := psr.Build()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var gb bytes.Buffer
	if *ty == md {
		gb = genMarkdown(doc)
	}

	if *ty == tx {
		gb = genText(doc)
	}

	err = ioutil.WriteFile(*fn, gb.Bytes(), 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

func walkFn(path string, info os.FileInfo, err error) error {
	if strings.HasPrefix(info.Name(), ".") {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if !info.Mode().IsRegular() {
		return nil
	}

	if !isSrcFile(filepath.Ext(path)) {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	strs := strings.Split(string(data), "\n")
	batches := trim.TrimDsDoc(strs)
	if len(batches) <= 0 {
		return nil
	}

	for _, bt := range batches {
		err = psr.Parse(bt, info.Name())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	return nil
}
