package main

import (
	"fmt"
)

func maina() {

	const a = 2 //aslında burada veri tipi belirsizdir, varsayılan bir değer olarak atanır
	//print edince, kullanılınca ..  // NOT: BU DURUM const için geçerlidir.
	println(a)
	//var b = 10.5
	//fmt.Printf("%T", a) // int
	fmt.Printf("%T %v \n", a*10.99, a*10.99)

	var c = 8 / 5          // int olarak düşündü go, ,var'da böyle
	const d = 8 / 5        // aynı üstteki gibi
	const e = 1000 * 10.99 // değişti float64 olarak gösterdi sonucu
	var f = 100 * 9.9199   // değişti float64 olarak gösterdi sonucu
	var g = int64(f)       // type conversion
	println(c, d, e, f, g)

	var F float64 = -40.18
	var K float64 = 0
	K = (F-32)/1.8 + 273

	fmt.Printf("\n %.2f", K) // yuvarladık
}
