package main

func main() {

	var x = true

	if x == false {
		println("FALSE")
	} else {
		println("true")
	}
	/////////////////
	if true {
		println("Öylesine yazıldı, sonuçta true kabul edildi")
	}
	/////////////////
	if 17%2 == 0 {
		println("dgdfg")
	}
	////////////////
	var a = 32
	if a > 32 {
		println("if blok")
	} else if a%2 == 0 {
		println("else if ilk blok")
	} else if a == 32 {
		println("else if ikinci blok")
	} else {
		println("else blok")
	}
	////////////////////
	///////////part 2 if else
	z := 45
	if z := 15; z == 15 { // bu blok içindeki atanan z var
		println("bu if scope içindeki farklı z, 15 değerine atanıp koşullandı", z)
	}
	println("Scope dışındaki z", z)

	// böyle aynı isimde olması doğru değil idi tabiki

}
