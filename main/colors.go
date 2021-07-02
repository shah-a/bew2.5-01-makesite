package main

import "github.com/fatih/color"

/*
 * colors.go initializes text formats to be wrapped around
 * strings anywhere in `package main`. For example, I can
 * go to `main.go` or `makesite.go` and write:
 * fmt.Println(red("This is red."))
 * or
 * fmt.Println("This is " + bold("bold"))
 */

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var boldRed = color.New(color.FgRed, color.Bold).SprintFunc()
var boldGreen = color.New(color.FgGreen, color.Bold).SprintFunc()

// var bold = color.New(color.Bold).SprintFunc()
