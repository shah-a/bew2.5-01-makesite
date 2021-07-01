package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Data struct {
	Text string
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "filename to parse")
	flag.StringVar(&file, "f", "", "filename to parse")
	flag.Parse()

	if file == "" {
		panic("\"--file\" (or \"-f\") flag required")
	}

	file = strings.TrimSuffix(file, ".txt") // remove ".txt" extension from filename string

	text, err := ioutil.ReadFile(file + ".txt")
	if err != nil {
		panic(err)
	}
	context := Data{string(text)}

	path, err := os.Create(file + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(path, context)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully generated \"%s.html\"!\n", file)
}
