package main

import (
	// b "code/code/practicemodule/basicvars"
	// c "code/code/practicemodule/complexvars"
	// p "code/code/practicemodule/programflow"
	// r "code/code/practicemodule/reusablecode"
	f "code/code/practicemodule/filesnweb"
	"fmt"
)

func main() {

	x := f.GetApi()
	fmt.Println(x)	
	// result := f.AgifyFromJson(x)
	// fmt.Println(result)	
	// for _, i := range result {
	// 	fmt.Println(i)
	// }
}