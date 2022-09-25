package reusablecode

import "fmt"

//this is equivalent to a void function since it doesn't return anything
func DoSomething() {
	fmt.Println("I do something")
}

//a funciton that returns data
func sum(a int, b int) int {
	total := a+b
	return total 
}

/*
above we explicitally state the type of both a and b but we can equally say
fun sum(a, b int) ... and go will be ok with this.
*/

//now let's return multiple values; we have to include the return type of the values we want to return
func AddVals(a, b int) (int, int) {
	anArray := []int{a,b}
	total := 0

	for _, i:= range anArray {
		total += i
	}

	length := len(anArray)
	return total, length 
}