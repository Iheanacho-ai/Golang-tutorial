package main

import "fmt"

func main() {
	age:= 35
	name:= "amarachi"

	//print
	fmt.Print("hello,")
	fmt.Print("world! \n") //new line
	fmt.Print("new line \n")

	//Println
	fmt.Println("hello world")
	fmt.Println("goodbye world")

	fmt.Println("my name is", name, "and i am", age)

	//Printf(formatted strings)
	fmt.Printf("my age is %v and my name is %v", age, name)
	fmt.Printf("my age is %q and my name is %q", age, name) // puts quotation over the printed out string
	fmt.Printf("age is of type %T \n", age)// prints out the type of the age variable
	fmt.Printf("you scored %f points! \n", 225.55) //prints out the text with the float number
	fmt.Printf("you scored %0.1f points! \n", 225.55) // prints out the text with the float number but rounds it up to one decimal points

	//Sprintf( save formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println("the saved string is:", str)

}