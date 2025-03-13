package main // this is important to be named "main" package, because "main" is a special package name as it tells Go this is the main entry point of the program

import "fmt" // this is a package (from standard library)

// this function is called "main" because go will execute this function when the program starts
func main() {
	fmt.Print("Hello World")
}

// run: go run app.go
// we can use go build even without go being installed
// remember a module consists of many packages, a module can be considered as a project
// in order to create a module, we can use go mod init
// example: go mod init example.com/first-app
// in windows, it will give u an exe file (first-app.exe)
// in linux, it will give u a binary file (first-app)
// in mac, it will give u a binary file (first-app)

// use go mod init to create a module
//example: go mod init example.com/first-app
// use go build to build a module after creating a module
// in windows, we can start the module by running the exe file using the command: ./first-app.exe
