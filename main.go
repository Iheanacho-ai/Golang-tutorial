package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
		
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader) // tip is the number the user inputs
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("you chose to save bill", b)
	default:
		fmt.Println("oops, that was not a valid option")
		promptOptions(b) // reruns the functio, and asks them to choose an option again
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

}


