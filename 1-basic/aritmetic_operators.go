package main

import "fmt"

func main() {
	z := 5 / 2 // integer verir
	fmt.Printf("%T %v \n", z, z)

	z2 := float64(5 / 2)
	fmt.Printf("%T %v \n", z2, z2)

	x := 1
	x++
	fmt.Println(x + 1)
	//fmt.Println(x++) // x++ go'da bir statement 'dır. iki statement ifade bir satırda olmaz. burada
	// ilk statement print. ikincisi ise x++ . x:=1 gibidir x++.

}
