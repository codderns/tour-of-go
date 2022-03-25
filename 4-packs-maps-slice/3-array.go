package main

import "fmt"

func main() {

	array1 := [...]int{1, 23, 4, 5, 6, 4, 8, 9}

	array2 := [...]string{"f", "f", "fdf", "fdg", "e", "f", "e", "fdg"}

	for _, v := range array1 {
		print(v)
	}
	println()
	fmt.Println(array2)
	////////////////////////

	var dizi [5]int
	dizi[0] = 15      // changed first indis
	fmt.Println(dizi) // we see confined array with 5 digit . [0 0 0 0 0]

	var dizi2 [5]int
	fmt.Println(dizi == dizi2) //print false, because we changed first index of dizi

	for i := 0; i < len(dizi); i++ {
		print(dizi[i], " - ")
	}
	println()
	///////////////////////
	newarr := [5]int{10, 100, 20, 200, 1221}

	fmt.Println(square(newarr))

}

func square(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
