package main

import (
	"fmt"
	"time"
)

func main() {

	//let's create some time variables
	n := time.Now()
	fmt.Println(n)

	//we can set the time to a specific moment too
	t := time.Date(2009, time.November, 12, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	//and we can parse the time from a string
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The parsed time is %s\n", parsedTime)
}
