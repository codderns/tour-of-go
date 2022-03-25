package main

func main() {

	var another_scope_this_var = 100000

	switch grade := 4; grade {
	case 5:
		println("perfect")
		ov := "biravvo" // sadece bu scope içinde yani "case 5:" içinde geçerlidir bu değişken
		println(ov)
	case 4:
		println("good")
	case 3:
		println("not bad")
	case 2:
		println("too close")
	case 1:
		println("shit")
	default:
		println("only digits between 1-5")
	}

	println(another_scope_this_var)

	switch {
	case false:
		println("This inf cannot be printed in case false")
	case true:
		println(true)
	}
}
