package main

import (
	"fmt"
)

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee  // taken some feature from employee
	hasDegree bool
	///  Ancak kesinlikle birbirlerinden bağımsızdırlar. Sadece oradaki field'ları almış.
}

// struct like above is called "underlying type"
// types like employee are called "defined type","named type"

func main() {
	e1 := employee{
		name:      "Serah",
		age:       21,
		isMarried: false,
	}
	fmt.Println(e1) //{Serah 21 false}
	e2 := e1        //assign to e2
	fmt.Println(e2) //{Serah 21 false}

	e2.name = "Bella"
	fmt.Println(e1, e2) // {Serah 21 false} {Bella 21 false} . So, discrete each other anymore

	// create a manager
	m1 := manager{
		employee: employee{ //We took these fields in the code from that structure.
			name:      "Karl",
			age:       45,
			isMarried: true,
		},
		hasDegree: true,
	}
	fmt.Printf("%#v \n", m1) //main.manager{employee:main.employee{name:"Karl", age:45, isMarried:true}, hasDegree:true}
	fmt.Println(m1)          //{{Karl 45 true} true}

	m1.name = "New-Karl"
	m1.age = 50
	m1.hasDegree = false
	fmt.Println(m1) // {{New-Karl 50 true} false}

	/// anonym struct
	theBoss := struct { // if a struct will be needed just one, then we can use it.
		//So, we cannot produce struct variables from anonym struct
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss) // {THE BOSS true}
}
