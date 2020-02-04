package main

import "fmt"

func tester() { //only called by Go runtime

	var hello = "Hello"
	fmt.Println("Uwaifo Idehenre")
	fmt.Println("Medua Osamo")
	fmt.Print("Print")
	fmt.Print("Print again")
	fmt.Println(hello)
}

//To run multiple go file : go run *.go
//You can also call it like this: `go run .` (notice the dot at the end).
// Or: Assuming there are file1.go, file2.go, and file3.go in the same directory: `go run file1.go file2.go file3.go

//Go build can compile both executable and library packages

//??? An example of a lirary package is "package lib"

//Declared names within a scope are unique and cant be duplicated
