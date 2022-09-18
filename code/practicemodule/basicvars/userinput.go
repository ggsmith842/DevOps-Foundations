package basicvars

import (
	//we can import multiple packages like so 
	"fmt"
	"bufio"
	"os"

)

func UserInput() {

	//create a new reader with an input parameter
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Text: ")
	//let's define multiple variables
	input, _ := reader.ReadString('\n') //the \n tell the reader to stop after a newline

	fmt.Println("You entered: ", input)

}