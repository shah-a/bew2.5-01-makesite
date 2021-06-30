package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}
