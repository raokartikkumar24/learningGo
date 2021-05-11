//offset is random
//maps are passed by reference, cheap to pass
//unsafe for concurrency
package main

import (
	"fmt"
)

func main() {

	IPLLeagues := make(map[string]int)
	IPLLeagues["RCB"] = 8
	IPLLeagues["CSK"] = 2
	IPLLeagues["MI"] = 1

	fmt.Printf("\n IPL leagues = %v", IPLLeagues)

	IPLTitles := map[string]int{
		"RCB": 0,
		"MI":  5,
		"CSK": 4,
	}

	fmt.Printf("\n IPL Titles = %v", IPLTitles)

	//Iterating the maps
	//map is unordered list. The offset is very randomised.
	//The output here will varry everytime we run the code
	for key, val := range IPLTitles {
		fmt.Printf("\nKey = %v \t Value = %v", key, val)
	}

	//update the values at key
	IPLTitles["MI"] = 6
	fmt.Println("\n", IPLTitles)

	//Add a new key to the map
	IPLTitles["KKR"] = 1
	fmt.Println("\n", IPLTitles)

	//delete the key from map
	delete(IPLTitles, "KKR")
	fmt.Println("\n", IPLTitles)
}
