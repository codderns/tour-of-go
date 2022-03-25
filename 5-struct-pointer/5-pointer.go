package main

import "fmt"

// func main() { // GO - > language structre -> pass by value

// 	x := 5

// 	fmt.Println(x) // 5
// 	// GO'da fonksiyon çalışınca aldığı parametre argümanın bir kopyasıdır
// 	double(x)      // 10
// 	fmt.Println(x) // 5
// 	// x value didn't change. GO is the language that pass by value

// }

// func double(num int) {
// 	num *= 2
// 	fmt.Println(num)
// }

func main() {
	myslice := []int{1, 10, 100, 1000}
	fmt.Println(myslice) // [1 10 100 1000]
	double(myslice)      // [2 20 200 2000]
	fmt.Println(myslice) // [2 20 200 2000]

	/*Why, because contrary to GO language structure, slice's language structre is "pass by reference"
	it causes a thing like above*/
	/*Otherwise, value wouldn't have changed*/

	/////////
	x := 5
	fmt.Println(x) // 5
	double2(&x)    // 0xc0000ae0f0   ve   10
	fmt.Println(x) // 10  // so, pointer caused the x value here

	// double has pointer method anymore

}

func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
}

// array için deneseydik yine değişmeyecekti, bir örneğini kopyala işle ve bas mantığı var normalde

func double2(num *int) { // this parametre is pointer anymore

	fmt.Println(num) // 0xc0000ae0f0
	// verinin adresindeki değer * ile yakalanır
	*num *= 2
	fmt.Println(*num) // 10

}
