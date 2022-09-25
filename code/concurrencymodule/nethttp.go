package main

import (
	"fmt"
	"net/http"
)

func RunHello() {
	//configure path and handler functoin
	http.HandleFunc("/hello",Hello)


	//Listen on port 8080 and block main
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080",nil)

}

//function we want to return in the API call
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}