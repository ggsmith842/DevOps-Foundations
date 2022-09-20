package complexvars

import "fmt"
import "sort"

func Slices() {

	/*
	Slices are abstractions built on top of arrays.
	Slices are resizable and can be sorted easily.
	*/
	//this is an array
	colors  := [3]string{"red","green","blue"}

	//this is a slice
	colors_slice  := []string{"black","purple"}
	colors_slice = append(colors_slice, "orange" )

	fmt.Println(colors, " is an array")
	fmt.Println(colors_slice, " is a slice")

	//if we want to only select certain elements we can using index slicing
	fmt.Println(colors_slice[1:]) //go from second element to the very end


	//you can preallocate the size and type of a slice with make
	numbers := make([]int,3) 
	numbers[0] = 123
	numbers[1] = 45
	numbers[2] = 27
	fmt.Println(numbers)

	//we can also sort slices with the sort package
	sort.Ints(numbers) 
	//notice all these actions are done inplace
	fmt.Println(numbers, " is a sorted slice!")

}