// https://pkg.go.dev/std - standard library

package main

import (
	"fmt"
)

func main() {
	// for loops
	x := 0

	//1
	for x < 5 {
		fmt.Println("vaue of x is:", x )
		x++
	} 

	//2

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	//3
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for in loop

	for index, value := range names{
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	//for in loop without either index or value

	for _, value := range names{
		fmt.Printf("the values are: %v \n", value)
	}
}