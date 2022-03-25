package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := first(31)
	if err == nil { // that's it. the first func returns two value(string, error).
		//so , we made a query based on two equivalents of the value.
		println("Your result", result)
	} else {
		fmt.Println(err)
	}
	////////////////////////

	result2, err2 := second(82)
	if err2 == nil {
		fmt.Printf("Your number's root is %.2f \n", result2)
	} else {
		fmt.Println(err2)
	}

}

func first(num int) (int, error) { // also it returns error with together string
	//return "fd", errors.New("sdfs")  // if bloğu içinde olmasa idi iki adet return'e gerek yoktu tabi ki

	if num%2 != 0 {
		return 0, errors.New("This number isn't even.")
	} else {
		return num, nil // nil, başlangıç değeri olmayan ifadelere verilen değerdir.
	}
	// ikinci parametre error olduğundan ve onu da girmek gerektiğinden nil olarak belirtilir
	// iki adet return belirtmek gerekir if bloğundan dolayı

}

func second(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negative numbers cannot be rooted.")
	} else {
		return math.Sqrt(num), nil
	}
}
