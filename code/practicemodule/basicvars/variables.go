package basicvars //declare the package this file belongs to

import "fmt" //import other packages;fmt is from the standard library

func Variables() {

	//declare and initialize a variable
	var aString string = "This is Go!"
	
	//we can also declare and initialize this way 
	aString2 := "This is also Go!"
	/*
	An important note is that := only works if we declare the 
	variable inside of a function.
	*/


	fmt.Println(aString)
	fmt.Println(aString2)

	//we can findout what the type of a variable is with PrintF
	fmt.Printf("The variable's type is %T\n", aString)

	//let's declare an integer
	var anInteger int = 42
	fmt.Printf("anInterger is of type %T\n",anInteger)

}

//here is a variable we declare outside of a function
var noFunc string = "I'm not in a function :("

//we can also declare constants 
const aConst int = 64