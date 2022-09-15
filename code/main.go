package main

import(
	"fmt"

)

func AddSum(a int, b int) int {
	var myTotal int
	myTotal = a + b

	return myTotal
}


func main() {
	fmt.Println("Hello World")
	fmt.Println(AddSum(2, 3))
}