
package main

import (
	"fmt"
	"strings"
)

func getInitials (n string) (string, string){
	s := strings.ToUpper(n) // TIFA LOCKHART
	names := strings.Split(s, "") // [TIFA, LOCKHART]

	var initials []string // {}
	for _, v := range names {
		initials = append(initials, v[:1]) // appends the items in postion 0 of the array value currently worked on in the loop, through to position 1, does not include 1 though, to the initials array
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} 
	
	return initials[0], "_"
}


func main() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)
}

