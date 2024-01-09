// https://pkg.go.dev/std - standard library

package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello")) // searches if hello is in the greeting variable

	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // replaces hello with hi in the greetings variables

	fmt.Println("original string value =", greeting) // shows that replace does not alter the original string

	fmt.Println(strings.ToUpper(greeting)) // to uppercase

	fmt.Println(strings.Index(greeting, "ll")) // shows the position of "ll" in the greetings string

	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // sorts the slice(flexible array), according to the largest number
	fmt.Println(ages) // alters the actualslice


	index := sort.SearchInts(ages, 30) // finds the number in the slice
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names) // sorts the strings in alphabetical order
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}