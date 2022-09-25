package filesnweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://api.agify.io?name=Michael"


func GetApi() string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n",resp)

	defer resp.Body.Close()


	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)

	return content
}


//now we need a way to work with the response
type AgifyResp struct {
	Name string
	Age, Count int
	
}

func AgifyFromJson(content string) []AgifyResp {

	person := make([]AgifyResp,0,20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_,err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var entry AgifyResp
	for decoder.More() {
		err := decoder.Decode(&entry)
		if err != nil {
			panic(err)
		}
	}

	person = append(person, entry)

	return person
}