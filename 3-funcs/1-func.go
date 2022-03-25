package main

import "fmt"

func main() {
	a := 1
	print(first(a), "\n") // it must be same type with func parameter
	//////////////

	var x = "My name is "
	var y = "Codder"
	println(second(x, y))
	/////////////

	fmt.Printf("%.2f", third(5.9, 4)) // rounded the float. Also parameter order is important
	///////////////

	fourth("\nThis", "function")
	////////////////

	fifth(45, "Anonym")
}

/// funcs ////
func first(value int) int { // what types in bracket and what return out of bracket
	return value
}

func second(a, b string) string {
	return a + b
}

func third(c float64, d int) float64 {
	return float64(d) * c * 2
}

func fourth(e, f string) {
	print(e + " " + f + " is printed")
}

func fifth(g int, h string) {
	fmt.Printf("\nMy name is %s and i am %d", h, g)
}
