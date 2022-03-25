package main

// What is the package

// A collection of pieces of code that do similar work is called a package.

// https://pkg.go.dev/std

import ( // for instance, import contents is a package
	"fmt"
	"strings"
)

func main() {
	// we must use package if imported

	fmt.Println("dfsdf")

	var s string = "Independent"
	fmt.Println(strings.Contains(s, "dent")) // returns true
	fmt.Println(strings.Contains(s, "tend")) // returns false
	fmt.Println(strings.Contains(s, ""))     // returns true

	fmt.Println(strings.Count(s, "e")) // returns 3. So, it shows how many e words it contains

	fmt.Println(strings.ToUpper("my name is fiction ğı"))

}
