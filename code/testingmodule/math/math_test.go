package math
import (
	 "testing"
)


/*
To run tests in go, the file must end with `_test.go`.

Tools like vs code will give you an interative option to run the test 
but you can also run the command `go test`. 

Then go will run the test for every file that ends with `_test.go`.
*/

//example using table driven testing.
func TestAddition(t *testing.T) {

	//this struc contains out table of test cases
	cases := []struct {
		desc	string
		a,b		int
		expected	int
	} {
		{"TestBothZero",0,0,0},
		{"TestBothPositive",2,5,7},
		{"TestBothNegative",-2,-5,-7},
	}

	//this for loop runs the test for each case in the table
	for _, testCases := range cases {
		actual := Addition(testCases.a,testCases.b)
		if actual != testCases.expected {
			t.Fatalf("%s: expected: %d got: %d for a: %d and b :%d",
			testCases.desc, actual, testCases.expected,testCases.a,testCases.b)
		}
	}

}

/*
Table-driven testing allows for much more concise code.

Instead of individually test each case in its own function call,
we are abble to run all cases inside one for loop and see the results. 
*/