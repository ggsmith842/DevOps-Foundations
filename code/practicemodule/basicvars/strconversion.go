package basicvars

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func StrConversion() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	//the _ char is an error object we are ignoring for now
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)

	fmt.Print("Now what's your favorite number? ")
	inputNum, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(inputNum),64)

	//here if we have an error it gets printed. nil is similar but not exactly null
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your number is ", aFloat)
	}

}