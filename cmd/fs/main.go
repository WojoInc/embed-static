package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed templates
var embedTemplates embed.FS

type templateData struct {
	// Destination package to place file in
	Package string
	// Name of file to create
	Filename string
	// Path to directory that should be embedded
	Directory string
	// Path to file which should serve as the index
	// i.e. /index.html
	IndexFile string
}

func parseInput() *templateData {
	if len(os.Args) != 5 {
		panic(fmt.Errorf("expected exactly 4 arguments"))
	}
	data := templateData{
		Package:   os.Args[1],
		Filename:  os.Args[2],
		Directory: os.Args[3],
		IndexFile: os.Args[4],
	}
	// If user specifies path relative to the directory to embed, remove
	// this prefix from IndexFile
	if strings.HasPrefix(data.IndexFile, data.Directory) {
		data.IndexFile = strings.TrimPrefix(data.IndexFile, data.Directory)
	}

	// Check the

	return &data
}

func main() {
	tmpl, err := template.ParseFS(embedTemplates, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	data := parseInput()

	// If the directory for the destination file does not exist, create it
	// Directory should already exist, because files to be embedded must be in a
	// (same/sub) directory of the go file containing the embed directive
	// This simply covers the case where `go generate` is executed before the
	// embedded files are created.
	dir := path.Dir(data.Filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0750); err != nil {
			panic(err)
		}
	}
	f, err := os.Create(data.Filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	temp := bytes.Buffer{}
	err = tmpl.Execute(&temp, data)
	if err != nil {
		panic(err)
	}

	if _, err = f.Write(temp.Bytes()); err != nil {
		panic(err)
	}
}
