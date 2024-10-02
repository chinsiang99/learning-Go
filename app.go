package main // this is important to be named "main" package, because "main" is a special package name as it tells Go this is the main entry point of the program

import "fmt" // this is a package

func main() {
	fmt.Print("Hello World")
}

// run: go run app.go
// we can use go build even without go being installed
// remember a module consists of many packages, a module can be considered as a project
// in order to create a module, we can use go mod init
// example: go mod init example.com/first-app
// in windows, it will give u an exe file (first-app.exe)
