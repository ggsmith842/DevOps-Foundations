package complexvars

import "fmt"
/*
A struct in GO is a custom data type. It is most similar to Java and Python classes.
There is no inheritance though like in Java.
*/

/*
Naming the Dog with an uppercase means it will be exported. If we use lowercase this is 
like saying its private. The same is true for functions. 
*/

// Dog is a struct
type Dog struct {
	//define attributes
	Breed string
	Weight int

}

func MakeDog() {
	Max := Dog{
		Breed: "Border Collie",
		Weight: 45	}


	//different ways to view attributes
	fmt.Println(Max.Breed)

	fmt.Printf("%+v\n", Max)

	fmt.Printf("Breed: %v\nWeight: 5%v\n", Max.Breed,Max.Weight)

	//we can also set values like so
	Max.Weight = 50	
	fmt.Printf("%+v\n", Max)
}