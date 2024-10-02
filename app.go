package main // this is important to be named "main" package, because "main" is a special package name as it tells Go this is the main entry point of the program

import "fmt" // this is a package

func main() {
	fmt.Print("Hello World")
}

// run: go run app.go
