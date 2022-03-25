package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%T %v \n", ("a" < "b"), "a" < "b") // bool true . true olma sebebi asci kod

	x := 1
	y := 2
	var z = 2
	var set1 = (x == y)
	set2 := (y == z)
	println(set1)                                          // bool bir değer olur set1 (false)
	fmt.Printf("%T %v \n", (set1 && set2), (set1 && set2)) // cevap bool false olur. false && true = false
	set3 := true
	println(set2 && set3) // true

}
