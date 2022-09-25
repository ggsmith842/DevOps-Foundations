package reusablecode

import "fmt"

//we can attach functions to struct to create methods
//recall we have the dog struct


type Dog struct {

	Breed string
	Weight int
	Sound string

}

//the first () is called the receiver. We also need a comment or we get a warning.
//Speak is how the dog speaks.
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

//make the dog speak x times
func (d Dog) ManySpeaks() {
	d.Sound = fmt.Sprintf("%v %v %v",d.Sound,d.Sound,d.Sound)
	fmt.Println(d.Sound)
}