package main

import (
	"fmt"
	"math"
)

// /*CHANNELS*/
// /**/

// // Önceki ders tekrarına önce bakalım
// type circle struct {
// 	r float64
// }

// var wg sync.WaitGroup

// func main() {

// 	wg.Add(1)
// 	var c1 = circle{25}

// 	fmt.Printf("%.2f \n", c1.area())

// 	go c1.display() // aynı anda farklı işler go routine ile olur. buraya yazınca main goroutine çalışacağı
// 	// sebebi ile çalıştırmazken waitgroup metodu bunu engeller
// 	wg.Wait()
// }

// func (c circle) display() {
// 	println("A Circle")
// 	wg.Done()
// }
// func (c circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// /// Channels

// type circle struct {
// 	r float64
// }

// func main() {

// 	var c1 = circle{25}

// 	//a1 := go c1.area() // hata olur go koymak. GO'da goroutine oluşturulunca ilişkili fonksiyonda
// 	// herhangi bir return değeri alınamaz. !!!!
// 	// Programın ana ilerleyişini asla bir return değerinin geri dönüşüne bağlayamayız.

// 	// Burada math.Pi * c.r * c.r var. Halledilir. Ama farklı bir kaynaktan bilinmeyen birşey alınınca
// 	// ne olur. GO garantici bir dildir, bu durumda goroutine yöntemi çalıştırılamaz.

// 	a1 := c1.area()
// 	go c1.display()
// 	fmt.Printf("%.2f \n", a1)

// }

// func (c circle) display() {
// 	println("A Circle")

// }
// func (c circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

/// İŞTE CHANNEL'ların olma sebebi de budur. Goroutine'lerin birbirleri ile iletişime geçmesi
// için channels kullanılır.
// Channels goroutine'lerin birbirleri ile iletişime geçmesini sağlayacağı gibi aynı zamanda
// bir goroutine tarafından gönderilen değerin diğer goroutine'de kullanılmaya başlamadan önce de
// gönderildiğinden emin olur.

/*goroutine'e return göndermede channel işe yarar*/

// channellar goroutine'lerin birbirleri ile iletişim kurmalarına da yarar
// func main() {

// 	// channel oluşturmak için anahtar kelime: chan
// 	// 	// ilk olarak identifier oluşturduk
// 	// 	var myChannel chan string // veri tipini de belirtmek lazım
// 	//
// 	// 	myChannel = make(chan string)
// 	// bunun yerine direk deklare edip atamasını yapabiliriz

// 	myChannel := make(chan string)

// 	go merhaba(myChannel) //parametre olarak bir channel alacak

// 	// channel'daki değer alınır
// 	fmt.Println(myChannel)   // 0xc00004a060 adresini gösterdi, değerini görmek için
// 	fmt.Println(<-myChannel) // <- operatorü kanaldaki değeri almaya işe yarar
// }

// // parametresi chanel olan merhaba fonks göndereceğiz
// func merhaba(chan1 chan string) {
// 	// return yerine aşağıdaki işlem yapılır
// 	chan1 <- "Merhaba" //değer olan string'i bir kanala gönderdik
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// func main() {
// 	chan1 := make(chan string)

// 	chan1 <- "merhaba"

// 	fmt.Println(<-chan1) //fatal error: all goroutines are asleep - deadlock!
// 	//burada hata var
// 	// bir channel bulunduğu goroutine'den başka bir goroutine'e değer dönmeden aynı channel
// 	// vasıtası ile kullanıldığı goroutine'i bloklar.
//	//Bu çook önemli bir meseledir
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

type circle struct {
	r float64
}

func main() {

	var c1 = circle{25}
	chan1 := make(chan float64)

	// return edebiliriz channel üzerinden.
	go area(c1, chan1)
	fmt.Printf("%.f2 \n", <-chan1) // istediğimiz değer bir kanalda //19632

	c1.display() // A Circle

}

func (c circle) display() {
	println("A Circle")

} // ilk parametre çember struct'ı, ikincisi channel
func area(c circle, mychannel chan float64) {
	sonuc := math.Pi * c.r * c.r
	mychannel <- sonuc // sonuc'u kanala gönderdik
}
