package complexvars
import "fmt"

/*
new()
allocates but does not initialize memory
results in zeroed storage but returns a memory address

Using make() allocates and initializes memory. 
Make() allocates non-zeroed storage and returns a memory address
*/

func MemoryAlloc() {

	/*
	this is an invalid operation.
	this will cause the program to crash
	*/
	// m := new(map[string]int)
	// m["theAnswer"] = 42
	// fmt.Println(m)

	/*
	If your intent is to immediatly add data to an object
	you must use make(). 
	*/
	m := make(map[string]int)
	m["theAnswer"] = 42
	fmt.Println(m)
	
}