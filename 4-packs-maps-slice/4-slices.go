package main

import "fmt"

func main() {

	// the slices is limitless arrays

	myslice := []int{}

	myslice = append(myslice, 10, 1, 2, 6, 6, 4, 5, 9) // to add the slice
	fmt.Println(myslice)

	/////
	myslice2 := make([]int, 4) //first size but resizeable, alterable -> [0 0 0 0] firstly
	myslice2[0] = 45           //of course, if there is at least 1 slice index. otherwise, use append
	fmt.Println(myslice2)

	///////
	myslice3 := make([]int, 2) //[0 0]
	myslice4 := []int{1, 9}

	myslice5 := myslice3
	myslice4[1] = 9856

	fmt.Println(myslice3, myslice5, myslice4)
	////////////////
	// slice slicing

	var slice6 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice6[:5])  // [1 2 3 4 5]
	fmt.Println(slice6[2:])  // [3 4 5 6 7]
	fmt.Println(slice6[2:5]) // [3 4 5]

	slice6 = append(slice6, myslice4...) // for solving this, we add three dot to end of variable that we want to change
	// we cannot add slice to slice, above method is to solve. This cause the elements of the array to be fragmented.
	fmt.Println(slice6) // [1 2 3 4 5 6 7 1 9856]

	/////////////////////////////////////////////////////// delete element of slices
	/// to  delete first n elements
	// slice6 = slice6[5:]
	// fmt.Println(slice6) // [6 7 1 9856]

	/// delete last n elements
	slice6 = slice6[:len(slice6)-3] // last  3 elements
	fmt.Println(slice6)             //[1 2 3 4 5 6]

	//
	lastslice := make([]bool, 2)
	fmt.Println(lastslice) // false false

}
