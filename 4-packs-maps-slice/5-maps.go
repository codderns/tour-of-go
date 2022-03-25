package main

import (
	"fmt"
)

func main() {

	// MAPS include key and value
	//            key   value
	mymap := map[string]int{
		"a": 2,
		"b": 9,
		"c": 7,
	}
	fmt.Println(mymap) //map[a:2 b:9 c:7]

	fmt.Println(mymap["b"]) // 9

	fmt.Println(mymap["ffferfe"]) // 0 . because there is not any key for "ffferfe" .

	value, isThere := mymap["FFFF"] //0 false  . So, if it is "a" then returns 1 true
	fmt.Println(value, isThere)

	////////////// // MAP element DELETE

	delete(mymap, "c")

	fmt.Println(mymap) // map[a:2 b:9]

	////////////////  MAP element ADD

	mymap["d"] = 122
	mymap["e"] = 355
	mymap["f"] = 477 //map[a:2 b:9 d:122 e:355 f:477]

	///// to assign the map to a variable

	assigned := mymap
	fmt.Println(assigned) //map[a:2 b:9]

	// if we delete any element from mymap then this affects the other variable that depends on.
	delete(mymap, "f")
	fmt.Println(mymap)    //map[a:2 b:9 d:122 e:355]
	fmt.Println(assigned) // map[a:2 b:9 d:122 e:355]

	// if added, same situation happens

	mymap["g"] = 45
	fmt.Println(mymap, assigned) //map[a:2 b:9 d:122 e:355 g:45] map[a:2 b:9 d:122 e:355 g:45] ->same

	// if we change assigned, mymap will be affected too
	delete(assigned, "g")
	fmt.Println(mymap, assigned) //map[a:2 b:9 d:122 e:355] map[a:2 b:9 d:122 e:355]

	///// to use for  loop
	for _, v := range mymap { // _ means dont print keys. if we dont want to print keys, we use _
		print(v, " ")
	}
}
