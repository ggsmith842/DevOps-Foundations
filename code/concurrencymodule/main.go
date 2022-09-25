package main


import (
	"fmt"
	"sync"
)

func main() {

	//initialize waitgroup
	var wg sync.WaitGroup

	//we are only waiting for one goroutine
	wg.Add(1)
	go hello(&wg)
	//this will block until the inner counter of the wait gorup is 0
	wg.Wait()
	goodbye()
}


/*
the & and * passes the WaitGroup using a pointer. This prevents a copy
waitgorup from being passed and a -1 panic occuring.
*/

func hello(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("hello")
}

func goodbye() {
	fmt.Println("goodbye")
}