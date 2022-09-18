package main

import (
	c "code/code/mymodule/mypackage"
	"fmt"
	
)

// main function to execute code
func main() {
	fmt.Println("Hello World")
	c.PrintHello()

}

//to run main we use go run main.go
//to package and distribute we need go build main.go
