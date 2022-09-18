package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	c "challengemodule/challengeutils"
	
)


func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 1st Value: ")
	input1,_ := reader.ReadString('\n')
	aNum1, err1 := strconv.ParseFloat(strings.TrimSpace(input1),64)

	fmt.Print("Enter 2nd Value:")
	input2,_ := reader.ReadString('\n')
	aNum2, err2 := strconv.ParseFloat(strings.TrimSpace(input2),64)

	
	if (err1 != nil) || (err2 != nil) {
		fmt.Println("An error occurred :/")

	} else {

		result := c.Add(aNum1,aNum2)
		fmt.Println("Total:", result)
	}

	

	

}