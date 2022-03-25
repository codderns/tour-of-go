package main

import (
	"fmt"
	"math"
)

/*
Go'da, interface kullanmak istersek, bizim için önemli olan aslında metotlardır.
Metotların yaptığı işlemler interface'i bağlamaz.
Örneğin, bir bilgisayar oyunu oynarken, gaming bilgisayar mı, normal masaüstü bilgisayar mı,
yoksa normal bir notebook mu olduğu farketmez. Ama sonuçta hepsi bir bilgisayardır.
*/
var a float64 = 3
var b float64 = 4

type rectangle struct {
	a float64
	b float64
}

// buradaki gibi bir interface oluşturup bu interface içerisine area(alan) ve circumference(çevre)
// metodunu yerleştirdiğimiz zaman, biz bu alan ve çevre metodunu interface üzerinden çalıştırabiliriz.
// Artık aldığımız parametreye göre de o alan ve çevre metodunun diktörtgene mi daireye mi ait olduğunu
// çıkartabiliriz.
type shape interface { /*interface yalnızca o fonksiyonların imzasına sahiptir (signature).
	 */
	area() float64
	circumference() float64 /*burada circumference adda imza var. o fonksiyonun nasıl çalışılacağı
	bilinmez, yalnızca metotlar ile ilgilenir. bunun işlemi parametre olarak alınan veri tipi yapar.
	buradaki işlemi aşağıdaki bunu veri tipi olarak alan fonksiyon yapıyor*/
}

// interface'e bağlı olan bir fonksiyon
func interfaceFunc(i shape) {
	fmt.Println(i)                 // alınan parametreyi göster	// {6 8}
	fmt.Println(i.area())          // area çalış				// 48
	fmt.Println(i.circumference()) // 28
	fmt.Printf("%T", i)            // main.rectangle
	fmt.Println()
}

func (r rectangle) area() float64 { // asıl uygulamayı rectangle struct'ından oluşturulmuş
	// bir fonksiyonun kendine ait olan area fonksiyonu yapacaktır, interface'den geldi
	return r.b * r.a

}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

////
type circle struct {
	r float64 // r is not relation with above // it represents the radius of circle
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func main() {

	r1 := rectangle{6, 8}                             // rectangle bir struct idi
	fmt.Println("Area ", r1.area())                   // 48   -> 6*8
	fmt.Println("Circumference ", r1.circumference()) // 28  -> 2*(6+8)

	interfaceFunc(r1)

	c1 := circle{10}
	fmt.Println("Circle Area ", c1.area())                   // Circle Area  314.1592653589793
	fmt.Println("Circle Circumference ", c1.circumference()) //Circle Circumference  62.83185307179586

	interfaceFunc(c1) // ↓
	// it wil be when it be its turn
	/*
		{10}
		314.1592653589793
		62.83185307179586
		main.circle  // data type.
	*/
}

func area() {

}
