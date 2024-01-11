package main

import (
	"fmt"
)


func main() {
	mybill := newBill("mario's bill")
	// mybill.format() // we can call the format receiver function 
	mybill.addItem("onion soup", 4.50)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}


