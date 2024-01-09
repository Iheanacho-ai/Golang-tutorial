package main

import "fmt"

func main() {
	//arrays
	//var ages [3]int =  [3]int{20, 25, 30}
	//shorthand of array
	var ages = [3]int{20, 25, 35}

	//an even shorterhand
	names:= [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi" // changes the string in position 1 to luigi

	fmt.Println(ages, len(ages)) //prints out the ages array,a nd its length
	fmt.Println(names, len(names))

	//slices (use arrays under the hoood, but its flexible)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)// add an item to the scores array
	fmt.Println(scores, len(scores))

	//slices ranges
	rangeOne := names[1:3] //prints out the item in postion 1 of the array names, to position 3, does not include 3 though
	rangeTwo := names[2:] // from position two to the end
	rangeThree := names [:3] // from position 0 to position 3 but not including 3
	fmt.Println(rangeOne) 
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)



}