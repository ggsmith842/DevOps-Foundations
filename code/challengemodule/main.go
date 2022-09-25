package main

import 	c "code/code/challengemodule/challengeutils"
import "fmt"

func main() {
	
	run := "Y"	
	for run != "N"  {
		run = c.GetRun()

		val1:=c.GetInput()
		val2:=c.GetInput()
		op := c.GetOperation()

		switch op {
		case "+":
			fmt.Println("The sum is: ",c.Add(val1,val2))
		case "-":
			fmt.Println("The difference is: ",c.Sub(val1,val2))
		case "*":
			fmt.Println("The product is: ",c.Multiply(val1,val2))
		case "/":
			fmt.Println("The quotient is: ",c.Divide(val1,val2))
		default:
			fmt.Println("Hmm..I don't know that one")

		
		}
		run = c.GetRun()
		if run != "Y" {
			fmt.Println("Thanks for calculating! Come again :)")
			break
		}

	}
	

		



}