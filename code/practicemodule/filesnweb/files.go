package filesnweb

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


func WriteContent(fileName string) {
	content := "Hello, from Go!"


	file, err := os.Create(fileName)
	CheckErr(err)


	length, err := io.WriteString(file,content)
	CheckErr(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	
	//means wait until everything is done then run this command
	defer file.Close()


}

func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	fmt.Println("Text read from file: ",string(data))

}