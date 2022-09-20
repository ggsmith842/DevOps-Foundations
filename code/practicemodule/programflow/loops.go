package programflow

import "fmt"

func PrintColors() {

	colors := []string{"red", "blue", "white"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	//range goes through the entire length of the iterable
	//below uses range and for each. We can return the index and the item.
	for index, color := range colors {
		fmt.Println("Color: ",color)
		fmt.Println("Index: ",index)
	}
	//alternatively if we didn't care about the index we can just put "_"
	// for _, color := range...

	//we can implement while logic using a boolean value
	value := 1
	for value < 4 {
		fmt.Println(value)
		value++ 
	}
}