
package main

import (
	"fmt"
)



func main() {
	menu:= map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping through maps

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key types
	phonebook:= map[int]string{
		2453367834: "mario",
		4567234478: "luigi",
		2133456789: "peach", 
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2453367834])

	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}

	//update maps

	phonebook[2133456789] = "bowser"
	fmt.Println(phonebook);

	phonebook[4567234478] = "yoshi"
	fmt.Println(phonebook);

}


