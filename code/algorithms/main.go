package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//generate a random array
	x:=generateSlice(15)
	fmt.Println(x)

	//call quicksort algorithm
	fmt.Println(quickSort(x,0,len(x)-1))

}

/*
This file implements a quicksort algorithm using recursion. Quicksort is a 
divide and conquor algorithm that works well on average O(n*log(n)). In the 
worst case it has a time of O(n^2). Most golang implementations of sort use
an optimized version of quicksort. Ref: https://blog.boot.dev/golang/quick-sort-golang/

The algorithm consists of two components.
1. Divide using the partition functions.
- partition finds the pivot and moves everything to the correct side
2. Conquor where we actual do the sorting.
- the quicksort function itself is how we apply the recursion to solve the problem
*/


/* This function takes last element as pivot, places the pivot element at its 
correct position in sorted array, and places all smaller (smaller than pivot) 
to left of pivot and all greater elements to right of pivot
https://www.geeksforgeeks.org/quick-sort/ */
func partition(arr []int, low, high int) ([]int, int) {
	
	//set pivot as the high value of the array
	var pivot int = arr[high]
	//set starting point as the start of the array 
	var i int = low

	//loop through the array until you have reached the pivot point
	for j := low; j < high; j++ {
		
		//if the value j is less than the pivot point
		if arr[j] < pivot {
			//swap the low left value with the right value
			arr[i], arr[j] = arr[j], arr[i]
			//move to the next value
			i++
		}
	}
	/*
	once j has reached the position before the pivot point, the ith value
	is where the high value (pivot) belongs in the array. 
	*/
	arr[i], arr[high] = arr[high], arr[i]
	
	
	/*
	Now all elements  smaller than the pivot are in the correct place, the 
	quicksort function will look at i and use it for the next pivot position.
	Quicksort will look at the left and the right of the new pivot position 
	recusively. 
	*/
	return arr, i 


}

//quicksort will recusively call partition until all values have been sorted
func quickSort(arr []int, low, high int) []int {
	
	//if low is not equal or greater than high
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		//Look at everything to the left of the pivot
		arr = quickSort(arr, low, p-1)
		//Look at everything to the right of the pivot
		arr = quickSort(arr, p+1, high)
	}

	return arr
}


//Generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	var slice = make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	return slice
}

