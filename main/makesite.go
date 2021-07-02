package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	Text string
}

// parseFlags checks for `--file` and `--dir` flags
func parseFlags() (string, string, error) {
	var file, dir string
	flag.StringVar(&file, "file", "", "filename to parse")
	flag.StringVar(&file, "f", "", "short form of \"file\" flag")

	flag.StringVar(&dir, "dir", "", "directory to parse")
	flag.StringVar(&dir, "d", "", "short form of \"dir\" flag")

	flag.Parse()

	if file != "" && dir != "" {
		return file, dir, errors.New("Please provide flag for only \"file\" or \"dir\" (not both)")
	}

	if file == "" && dir == "" {
		return file, dir, errors.New("Please provide flag for either \"file\" or \"dir\"")
	}

	return file, dir, nil
}

// appendFilePath adds all files that end in `.txt` to the `paths` slice in `main()`
func appendFilePath(file string, paths *[]string) {
	ext := filepath.Ext(file)
	if ext != ".txt" {
		panic("File must be \".txt\"")
	}

	// removes ".txt" extension from `file` string, then appends to `*paths`
	// removing ".txt" will help later when we need to rewrite as `.html`
	*paths = append(*paths, strings.TrimSuffix(file, ".txt"))
}

// appendDirPaths calls appendFilePath for all files in `dir` that end in `.txt`
func appendDirPaths(dir string, pathsPtr *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".txt" {
			continue
		}
		fullPath := filepath.Join(dir, file.Name())
		appendFilePath(fullPath, pathsPtr)
	}
}

// generateHTML parses a .txt file located at `path` and generates HTML based on its content
func generateHTML(path string) {
	text, err := ioutil.ReadFile(path + ".txt")
	if err != nil {
		panic(err)
	}
	context := Data{string(text)}

	htmlPath, err := os.Create(path + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(htmlPath, context)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully generated \"%s.html\"!\n", path)
}
