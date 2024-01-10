
package main

import (
	"fmt"
)

func updateName(x *string){ // we dereference the pointer passed as ana rgument
	*x = "wedge" // we change the derefernced variable to a different string
}


func main() {
	name := "tifa"
	// updateName(name)

	//fmt.Println("memory address of name is:", &name)

	m := &name // stores the pointer we created by using '&', and the variable name
	fmt.Println("memory address is:", &name)
	fmt.Println("value at memory address:", *m) // *m finds the variable using the pointer &name, basically dereferences the pointer

	fmt.Println(name)
	updateName(m)

	fmt.Println(name)

}


