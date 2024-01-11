package main

import (
	"fmt"
)


func main() {
	mybill := newBill("mario's bill")
	// mybill.format() // we can call the format receiver function 
	fmt.Println(mybill.format())
}


