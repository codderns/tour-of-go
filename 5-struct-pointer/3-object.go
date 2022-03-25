package main

import "fmt"

type mile float64

type kilometer float64

func main() {

	var m1 mile

	m1 = 5.9
	fmt.Println(m1)                // 5.9
	fmt.Printf("%T %v \n", m1, m1) // main.mile 5.9

	fmt.Println(m1 + 5.9) //11.8 okay but ↓

	f1 := float64(10.5)

	//fmt.Println(f1 + m1) // this is mistake. Because type of f1 is float64 but type of m1 is mile !
	// So, you cannot add two distinct type but ↓
	fmt.Println(mile(f1) + m1) // this is avalaible for summing

	fmt.Printf("%T %v \n", (mile(f1) + m1), (mile(f1) + m1))       // main.mile 16.4
	fmt.Printf("%T %v \n", (f1 + float64(m1)), (f1 + float64(m1))) // float64 16.4

	///////////////////////

	k1 := kilometer(7.8)
	fmt.Println(k1)                // 7.8
	fmt.Printf("%T %v \n", k1, k1) // main.kilometer 7.8
	// mile and kilometer is different data types .

	//fmt.Println(m1 + k1) // this is error because of summing two types

	////////
	m4 := mile(50)
	k4 := toKilometer(m4) // we consist of value from object named toKilometer !!!!!!!11
	fmt.Println(k4)       // 80

	k5 := kilometer(40)
	m5 := toMile(k5)
	fmt.Println(m5) // 24.8

}

// to use
////////////like i int
func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
