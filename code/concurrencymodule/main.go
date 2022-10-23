package main


import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//initialize waitgroup
	// var wg sync.WaitGroup

	// //we are only waiting for one goroutine
	// wg.Add(1)
	// go hello(&wg)
	// //this will block until the inner counter of the wait gorup is 0
	// wg.Wait()
	// goodbye()
	Ranging()
}


/*
the & and * passes the WaitGroup using a pointer. This prevents a copy
waitgorup from being passed and a -1 panic occuring.
*/

func hello(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("hello")
	//print the time to see when the function is called
	fmt.Println(time.Now())
}

func goodbye() {
	fmt.Println("goodbye")
	//print the time to see when the function is called
	fmt.Println(time.Now())
}

var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

//export this func
func Ranging() {
	// create a channel
	ch := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")

	for greeting := range ch {
		// sleep and print
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received!", greeting, time.Now())
	}

}

//writes a greet to the given channel
func greet(ch chan<- string) {
	fmt.Println("Greeter Ready!")

	//greet
	for _,g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeter Completed!")
}
