package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
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
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "got error running the file: %s", err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()
	fmt.Fprintln(out, outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	// Parse the mardown file though blackfriday and bluemonday to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
