package programflow

import "fmt"
import "math/rand"
import "time"



func Ifmain() {

	//if statement
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Negative"
	} else if theAnswer == 0 {
		result = "Zero"
	} else if theAnswer > 0 {
		result = "Positive"
	}

	fmt.Println(result)


	//switch statements
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1

	fmt.Println("Day", dow)

	var result2 string

	switch dow {
	case 6:
		result2 = "It's Friday!"
	case 7:
		result2 = "It's Saturday!!"
	case 1:
		result2 = "It's Sunday!"
	default:
		result2 = "It's some other day"

	}
	fmt.Println(result2)


}