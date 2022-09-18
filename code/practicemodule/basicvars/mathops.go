package basicvars

import (
	"fmt"
	"math"
)

func Mathops() {
	fmt.Println("Here is some math!")

	i1,i2,i3 := 12,13,14
	intSum:= i1 + i2 + i3
	fmt.Println(intSum)

	/*
	Floating values in go work a little differently with precision.
	You'll get a really weird output. We can use the math package to
	help.
	*/
	f1, f2 := 2.0, 3.4

	fSum := f1 + f2
	fmt.Println("Float sum: ", math.Round(fSum))
	//the problem here is we lose the remainder
	//so let's try something else
	fmt.Println("Float sum: ", math.Round(fSum*100)/100)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi

	fmt.Printf("Circumference: %.2f\n", circumference)

}

