package complexvars

import "fmt"


func Arrays() {

	//create an array of 3 string values
	var colors [3]string

	//initialize colors array
	colors[0], colors[1], colors[2] = "Red", "Orange", "Yellow"

	fmt.Println("Arrays")
	// print all the array elemnts
	fmt.Println(colors)

	//get the individual elements
	fmt.Println(colors[1])

	//get the number of elements in an array
	fmt.Println("Number of colors: ", len(colors))

	//initialize an array when you declare it
	numbers := [5]int{1,2,3,4,5}

	fmt.Println(numbers)
}