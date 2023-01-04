package main

import "fmt"

func main(){
	// var age uint32 = 26
	// var isMarried bool = true
	// var name string = "Alperen"
	// var weight float32 = 72.5

	//---Biz yukardaki gibi her değişkene var yazmak yerine aşağıdaki yöntemle daha temiz bir kod yazabiliriz.
	//Aşağıdaki yazımda boy kilo verilerini diğer verilerden ayıracak bir satır boşluk bırakıyoruz. Bu yazılı bir kural olmasa da
	//Genel konsept açısından uygulanır. Birbiriyle ilişkili veriler belli grup halinde yazılır.---

	// var(
	// 	name string = "Alperen"
	// 	isMarried bool = true
	// 	age uint32 = 26

	// 	weight float32 = 72.5
	// 	height int = 172
	// )

	//Ama yukardaki yönteme kıyasla daha derli toplu bir yazım biçimi var.

	// var name, isMarried, age, height, weight = "Alperen" , true , 26 , 172 , 72.5

	name, isMarried, age, height, weight := "Alperen" , true , 26 , 172 , 72.5 // bu da en kısa ve değişken tanımlamada sık kullanacağımız yöntem.

	//Biz bir değişkeni atayıp ona değer vermediğimizde o değişken default olarak zero değerini alır. name değişkeninde değer vermeden yazdırırsak
	//boş bir çıktı alırız. Bu değer bir int için 0 dır. String ifade içinse "" boş değerdir. Bool için ise bu false olarak tanımlanır.

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)
}