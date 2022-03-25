// package main

// import (
// 	"fmt"
// 	"time"
// )

// // GOROUTINES

// /*Birden fazla routin'ler sayesinde birden fazla işin aynı zamanda yapılmasını sağlayabiliriz..
// Yani go'da eşzamanlı yapılan görevlerin her birine bir GOROUTINE diyebiliriz.
// Burada "eşzamanlı" kelimesinin karşılığı "councurancy" dir.
// */

// func main() {

// 	/*Bir programın en hızlı çalışma şekli baştan sona kadar sırayla çalışma şekli değildir.
// 	Neden
// 	Biz web sunucusundan veri aldığımızı düşünelim, 5 sn olduğunu. Kabul edilimeyecek uzunlukta
// 	ise fonksiyonumuz, diğer işlemler de bunu bekleyecektir.
// 	O yüzden sıralı çalıştırma yöntemi değildir programı çalıştırmanın en hızlı yolu

// 	Hızlı olanı çalıştırabiliriz.
// 	İşlem önceliği verebiliriz.

// 	Peki birden fazla işi eşzamanlı çalıştırmak istersek ? İşte bunun için "goroutine" vardır

// 	Fonksiyon önüne go yazılır. Go routine oluşturulur.
// 	*/

// 	go printX()
// 	fmt.Println()
// 	go printY()
// 	time.Sleep(time.Second) // main fonksiyonun çalışasını 1 sn durdur.

// 	// Çalıştırınca önce fmt.println sonra XXXXXXXXXXXXXXXXXXXXYYYYYYYYYYYYYYYYYYYY
// 	// Tekrar çalıştır yine XXXXXXXXXXXXXXXXXXXXYYYYYYYYYYYYYYYYYYYY
// 	// Tekrar çalıştırınca bu sefer YYYYYYYYYYYYYYYYYYYYXXXXXXXXXXXXXXXXXXXX
// 	// SONRA XXXXXXXXXXXXXXXXXXXXYYYYYYYYYYYYYYYYYYYY

// 	// yani şöyle oldu, önce ana goroutine main çalışır, sonra iki go routine oluşur, o sırada çalışırlar
// 	// program çıkacakken, time.sleep ile kısa bir an beklenir ve diğer go'lar işlerini tamamlar.
// 	// Sonra diğerleri işlerini tamamlar ve yazdırır. B
// 	// Ama bazen X'ler bazen Y'ler çalışır. Biz sırasını kontrol edemiyoruz. Kontrol "channel" ile olur.

// 	// Mesela X'lerden 1000 tane Y'den 10 tane bastırsak Y'leri önce verir muhtemelen. Garanti değil.
// 	// Ancak ne kadar veri basacağını bilemez program. Gelecek olan verinin boyutunun ne kadar olmadığını
// 	// bilmediği durumlar sıkıntılıdır. ayrıca kendi main fonksiyonuna sleep komutu ile programın
// 	// performansını düşürmüş oluruz.
// 	// Farklı fonksiyonlar ile bunu yapmak isteriz.

// 	//sleep yazmazsak çalıştırınca yalnızca çalışan kod fmt.Println()'dir.
// 	/*Sebebi biz main fonksiyonunu çağırdığımız anda fonksiyonuzun ana GOROUTINE'i çalışmış olur.
// 	Çalışan tüm go programlarında bizim ana goroutine'miz olan main goroutine'i çalıştırmış oluruz.
// 	Main goroutine çalışır, ikinci goroutine de vardır printX . Üçüncü goroutine oldu printY için.
// 	Ancak diğer goroutine'lerin çalışmaya başlamasını beklemeden programın sonuna gelinir. Main
// 	goroutine çalışır ve kapanır, diğerleri olmadan program sonlanır.
// 	Biz bu çıktı alamama sorununu şöyle yapabiliriz.
// 	Main fonksiyonunun çalışmasını bir süre durdurarak bunu başarırız. Time sleep kullanarak bu olur
// 	*/

// 	/* Concuracy yani eşzamanlıdan kasıt, bir işlem(goroutine) fazla zaman alıyorsa onun yerine başka
// 	işlemleri al şeklindedir. Aynı anda çalışmaya parallelizm denir.*/
// }

// func printX() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Print("X")
// 	}
// }
// func printY() {
// 	for i := 0; i < 2; i++ {
// 		fmt.Print("Y")
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

// biz her zaman x'lerin daha önce yazılmasını istiyoruz mesela :
// bunnun için wait group kullanılabilir
var wg sync.WaitGroup

func main() {

	//wg.Add(1) // bir adet goroutine var diye belirttik

	go printX()
	wg.Wait() // bu işlem tamamlandı. Yani kendinden önceki goroutine tamamlanana kadar bekle yazdık
	go printY()
	// X'lerin önce Y'lerin sonra çalıştırıldığını gördük.

}
func printX() {
	for i := 0; i < 200; i++ {
		fmt.Print("X")
	}
	wg.Done() // işlemin tamamlandığının sinyalini ana fonksiyona  göndeririz
}

func printY() {
	for i := 0; i < 2; i++ {
		fmt.Print("Y")
	}

}
