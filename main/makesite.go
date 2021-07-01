package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Data struct {
	Text string
}

func main() {
	text, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	context := Data{string(text)}

	path, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(path, context)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully generated \"first-post.html\"!")
}
