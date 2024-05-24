package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8">
			<title>Markdown Preview Tool</title>
		</head>
		<body>`

	footer = `</body>
		</html>`
)

func main() {
	filename := flag.String("file", "", "Markdown file preview")

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, "got errorr running the file: %v", err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	return saveHTML(outName, htmlData)
}
