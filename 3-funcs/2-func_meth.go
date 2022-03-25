package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	print(first(81))
	///////////

	// time
	var x int = 1
	//variable type is time
	var today time.Time = time.Now() // If the code we're using isn't just any code we've written before, if we got it, it's a package.
	fmt.Println(x, today)

	///////  bufio  and  os
	print("Please input your exam result:")
	reader := bufio.NewReader(os.Stdin) //NewReader assigns the value entered from the keyboard to the reader variable with the help of the Stdin method.
	enteredvalue, _ := reader.ReadString('\n')
	//first is real value, second is error value must be returned. but if it is _ we can skip err part
	fmt.Println("Your score is", enteredvalue)

	//

	////////// //////////////////
	q, r := dividing(105, 8)
	println(q, r) // print the values of quotient and remainder
}

func first(grade int) string {
	if grade > 80 {
		return "AA" // if there is return and then if condition is ok, func is completed.
		// code below cannot work
	} else {
		return "BB"
	}
}

///    parameter values                  return values
func dividing(dividend, divisor int) (quotient, remainder int) { //this is called as multiple return

	quotient = dividend / divisor
	remainder = dividend % divisor

	return quotient, remainder
}
