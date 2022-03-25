package main

import "fmt"

/*
Go'da interface kavramı polymorphism yani çok biçimlilik kavramı ile ilişkilidir.
*/

/*
interface, yapılmasını istediğimiz metotların onlar için bize bir işaret yani signature taşıyordu.

Normalde bir klastaki özelliği alıp iki ayrı klasta kullanıp o iki klasın da aldığı
ortak klastan alınan özellik olmasına rağmen çıktıda farklı sonuçlar elde etmekdir.
GO'da ise bunu class yerine struct kalıbı içinde düşüneceğiz.
*/

/*Avantajı metodun tekrar tekrar kullanılabilmesidir. Tek farkı gelen veri tipine bu metodun
farklı şekilde uygulanmasıdır.*/

type shape interface {
	area() float64 // area metodumuzun veri tipi float64 olsun
}

// bu interface'e bağlı olan bir metot yazalım
func printArea(shapes ...shape) { //parametre olarak shapes adında shape veri tipinde değer alsın
	// ... anlamı varyatik parametre olarak aldık anlamındadır. Yani kaç tane shape parametresi içinde
	// olacak bilmiyoruz anlamındadır.

	for index, shape := range shapes { // gelen her şeklin elemanını shape değişkeni olarak belirtir
		fmt.Println(index+1, ". Alanı ", shape.area()) // ilgili shape'in area metodunu yazdırsın
	}
}

type triangle struct {
	a float64
	h float64
}

// üçgen alanı bulma fonksiyonu
func (t triangle) area() float64 { // fonksiyonumuz triangle veri tipine bağlı. fonksiyonumuz area, dönüş tipi f64
	// area() üsteki interface ismi ile uyuşması gerek ayrıca veri tipi de farklı olunca interface çalışmaz
	return t.a * t.h / 2

}

/// şimdi de kare için veri tipi yapak
type square struct {
	a float64 // bize tek kenar yeterli
}

func (s square) area() float64 {
	return s.a * s.a
}

/// rectangle
type rectangle struct {
	a float64
	b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func main() {

	ucgen := triangle{30, 10}
	kare := square{10}
	dikdortgen := rectangle{10, 20}

	// printarea metodumuzun içine üstte struct olarak oluşturduğumuz değişkenleri argüman olarak
	// verebiliriz,
	printArea(ucgen, kare, dikdortgen) /*metotları içine vermiş olmamızın sebebi printarea
	fonksiyonu içine varyetik fonksiyonları kullanmamızdan dolayıdır.. */

	/* aşağıdaki gibi görürüz metot çalışırsa
	1 . Alanı  150
	2 . Alanı  100
	3 . Alanı  200
	*/
}
