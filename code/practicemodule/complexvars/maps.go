package complexvars

import "fmt"
import "sort"


/*
In go a map is an unordered collection of key-value pairs.
This is essentially a hash table. Key are typically strings.
In python this is most similar to a dictionary 
*/

func Mapvars() {

	//here we are saying the key is of type string and the value is also of type string
	states := make(map[string]string)

	fmt.Println("An empty map ", states)

	//let's add some data
	states["MS"] = "Mississippi"
	states["CA"] = "California"
	states["CO"] = "Colorado"
	fmt.Println("We added some data:\n",states)

	//we can access the data by passing the key to the map
	california := states["CA"]
	fmt.Println(california)

	//we can delete items with delete(map, "KEY")
	//we can add items just like we did above

	//to iterate through a map we can use a for loop
	for k,v := range states {
		fmt.Println(k," : ",v)
	}

	/*	
	But now we want to order our data
	*/

	//let's initialize a slice
	keys := make([]string, len(states))
	
	//now we fill out slice with this for loop
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	//sort the slice
	sort.Strings(keys)
	//output the state values based on the ordered keys
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}