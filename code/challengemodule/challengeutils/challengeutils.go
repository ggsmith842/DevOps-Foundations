package challengeutils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func GetRun() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Y to calculate or N to quit ")
	runner, _ := reader.ReadString('\n')

	return runner
}



func GetInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter value: ")
	input, _ := reader.ReadString('\n')
	aNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		return 0
	} else {
		return aNum
	}
}


func GetOperation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Operation (+, -, *, /")
	operation, _ := reader.ReadString('\n')

	op := strings.TrimSpace(operation)

	return op
}

func Add(val1 float64, val2 float64) float64 {
	var sum float64 = val1 + val2

	//fmt.Println(sum)
	return sum
}

func Sub(val1,val2 float64) float64 {
	diff := val1 - val2

	return diff
}

func Multiply(val1,val2 float64) float64 {
	product := val1 * val2

	return product
}

func Divide(val1, val2 float64) float64 {
	result := val1 / val2

	round_result := math.Round(result*100)/100

	return round_result
}


