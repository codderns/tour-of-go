package main

import "fmt"

func main() {

	// Bir pointer başka bir değişkenin adresini tutan değişkendir.
	// A pointer is a variable that holds the address of another variable.

	name := "Michael"

	fmt.Println(name)

	fmt.Println(&name) // 0xc000038230 -> this is the address of the place in the computer memory where the variable name is located.
	// & -> address operator

	x := 11

	fmt.Printf("%T %v \n", x, x)   //int 11
	fmt.Printf("%T %v \n", &x, &x) //*int 0xc0000140c0
	// if we put &(address format) in front of x then result is *int 0xc0000140c0
	// it means, if we add &, our variable is pointer.

	////////////

	//var y int = &x  // impossible because of int and int* are not same
	//fmt.Println(y)

	////////////////////////////////////////////

	y := 10
	fmt.Println(y)     // 10
	fmt.Println(&y)    // 0xc0000ae0a8 -> the address of ten
	fmt.Println(*(&y)) // 10 -> the duty of * is dereferrecing
	// oradaki adresi gerçek değerine döndürür

	//////////////////////////////////////////////////

	x1 := 20
	x2 := x1
	x1 = 10
	fmt.Println(x1, x2) //10 20

	// pointer works in that point. if we assign the x2 to x1 as reference, then once we change x1,
	// x2 will also change
	x3 := 40
	x4 := &x3 // x4 is pointer anymore
	x3 = 30
	*x4 = 50                                    // we must put *. because x4 value is a reference, and we convert to int value
	fmt.Printf("%T x3: %v x4: %v", x4, x3, *x4) // *int x3: 50 x4: 50

	// we see that when x3 is changed, x4 also changes
	// we see that when x4 is changed, x3 also changes

	///////////////////////////////////////

	x5 := [4]int{1, 10, 100, 1000} // the array passes by value. // array'ler değer paylaşır
	x6 := x5
	fmt.Println(x6)

	x6[0] = 3

	fmt.Println(x5, x6) // [1 10 100 1000] [3 10 100 1000] // so, arrays are independent

	///////////// but slice

	x7 := []int{2, 20, 200, 2000} // the slice pass by reference !!!
	x8 := x7

	x8[0] = 4
	fmt.Println(x7, x8) //[4 20 200 2000] [4 20 200 2000] //so, slices work dependent on each other.

	// pointer ile yapılır

	x9 := []int{9, 99, 999, 9999}
	x10 := x9
	//	x9[0] = 10
	fmt.Println(x9, x10) // it works like pointer.

}
