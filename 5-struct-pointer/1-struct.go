package main

import "fmt"

// map'lerde key ve value vardı ve bunların taşıdıkları değerler birbirlerinin aynısı olmak zorundadır

// but struct is not same with maps.
// farklı veri tiplerinden faydalanma durumu var ve bu ihtiyaç struct'ın olmasına neden oldu.

func main() {
	var employee struct { // bu struct veri tipimiz, alttaki üç verinin üstüne inşa edilecektir.
		name      string
		age       int
		isMarried bool
	}
	fmt.Println(employee)          // { 0 false} -> as default, "" 0 false
	fmt.Printf("%#v \n", employee) //struct{name string; age int; isMarried bool }{name:"", age:0, isMarried:false}

	// add value
	employee.name = "Angel"
	employee.age = 25
	employee.isMarried = true

	// to reach its data
	fmt.Println(employee.name)
	///////////////////////////////

	// We want a struct with many values in a field. For this, we use type instead of var.
	type person struct {
		name      string
		age       int
		isMarried bool
		kids      []string // is consisted of arrays
	}
	// let create variable from person
	var p1 person
	p1.age = 45
	p1.isMarried = true
	p1.name = "Henry"

	var p2 person
	p2.age = 34
	p2.isMarried = false
	p2.name = "Thiago"

	fmt.Printf("%#v \n", p2) // main.person{name:"Thiago", age:34, isMarried:false}

	e3 := person{
		name:      "Fredick",
		age:       17,
		isMarried: false,
		kids: []string{
			"Joe", "April", "Katie",
		},
	}
	fmt.Println(e3)          // {Fredick 17 false [Joe April Katie]}
	fmt.Printf("%#v \n", e3) //main.person{name:"Fredick", age:17, isMarried:false, kids:[]string{"Joe", "April", "Katie"}}
	fmt.Println(e3.kids[0])  // Joe
}
