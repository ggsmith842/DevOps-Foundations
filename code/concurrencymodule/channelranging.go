package main

// import (
// 	"fmt"
// 	"time"
// )

// var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

// //export this func
// func Ranging() {
// 	// create a channel
// 	ch := make(chan string, 1)
// 	// start the greeter to provide a greeting
// 	go greet(ch)
// 	// sleep for a long time
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main ready!")
// 	for {
// 		// receive greeting
// 		greeting := <-ch
// 		// sleep and print
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println("Greeting received!", greeting)
// 	}

// }

// //writes a greet to the given channel
// func greet(ch chan<- string) {
// 	fmt.Println("Greeter Ready!")

// 	//greet
// 	for _,g := range greetings {
// 		ch <- g
// 	}
// 	fmt.Println("Greeter Completed!")
// }

