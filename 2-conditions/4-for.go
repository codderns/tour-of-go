package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 2 {
		print(i, " ")
	}
	print("\n")
	///////////
	var i = 5
	for ; i < 10; i++ {
		print(i, " ")
	}
	/////////////
	for {
		fmt.Printf("i am %v and i am written till forever if somebody doesn't stop me \n", i)
		if i > 12 {
			break
		}
		i += 2
	}
	////////////////
	var k = 0
	for k < 5 {
		k++
		if k == 2 {
			continue
			//print("2 won't printed on terminal")
		}
		print(k, " ")
	}
}
