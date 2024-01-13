package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills using the bill struct type
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string { // receiver functions, the parenthesis before the function name, makes this function only accessible from a bill object
	fs := "Bill breakdown: \n"
	var total float64 = 0 // the total variable

	// list items
	for k, v := range b.items { // we are mapping through the bill object we are taking
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)

	return fs
}

//update tip

func (b *bill) updateTip(tip float64){ // when you pass a struct to a receiver function, go makes a copy, so when you try to change it, it only changes the copy, thats why you should make a copy of the pointer instead
	b.tip = tip // however go dereferences the pointer automatically
}

// add an item to the bill
func (b *bill) addItem (name string, price float64) {
	b.items[name] = price
}

// save a bill
func (b *bill) save(){
	data := []byte(b.format())
	err:= os.WriteFile("bills/"+b.name+".txt", data, 0644) // first argument - location, second argument - data, third argument - permission

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to file")
}