package main

import (
	"fmt"
	m "mymodule/mypackage"
)


// main function to execute code
func main() {
	fmt.Println("Hello World")
	m.PrintHello()

}

//to run main we use go run main.go
//to package and distribute we need go build main.go
