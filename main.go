package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n') 

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // reader - reads the info the user types in; Stdin -says standard import, meaning the terminal

	name, _:= getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _:= getInput("Choose option (a - add item, t - add tip, s - save a bill): ", reader)
	
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader) // getInput is the function we created earlier
		price, _ := getInput("Item price ($): ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println("Tip: ", tip )
	case "s":
		fmt.Println("you chose s")
	default:
		fmt.Println("oops, that was not a valid option")
		promptOptions(b) // reruns the functio, and asks them to choose an option again
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

}


