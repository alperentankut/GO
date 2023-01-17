package main

import "fmt"

var packVar = "pack scope" //Paket düzeyi
func main(){

	var funcScope = "func scope" //Blok düzeyi
	var packVar = "var scope" //Biz burda packVar ı print ettiğimizde bloktaki pack var yazdırılacaktır. Aynı isimde paket düzeyindeki değişken
	//yazdırılmayacak.

	fmt.Println(funcScope)
	fmt.Println(packVar)
}

// var name = "arin"
// name := "elis"
//fmt.Prinln(name);

//---Yukarda hatayla karşılacağız. Bunun sebebi ise ilk başta name değişkenine
//hem declaration hem assignment yapıyoruz. Onun altında ise name değişkenine
//shorthand declaration yöntemiyle tekrar declaration ve assignment yapmaya çalıştığımız için hata alıyoruz.
//biz shortand declaration yerine name = "elis" şeklinde yazarsak hata almayız
//çünkü bu yazımda sadece assign etmiş oluyoruz. Fakat shorthand declaration yöntemiyle
// örneğin name , surname := "elis" ,  "software" şeklinde bir atama yaparsak
//yazdırırken hata almayız. Biz bir değişkeni iki kere declare edemeyiz.

//Biz fonksiyon scope u içinde oluşturduğumuz değişkenlere sadece fonksiyon içinde ulaşabiliriz. Fonksiyon dışında ulaşamayız.
//Fakat global scope da oluşturduğumuz değişkenlere fonksiyon içinde de dışında da ulaşabiliriz.
//Biz fonksiyon dışında hep anahtar kelimeleri kullanırız (import , var , func) bunları kullanmadan short declaration yaparsak (:=) hata alırız.
//çünkü biz öncesinde anahtar kelime belirtmediğimiz için ne yapmak istediğimiz anlaşılmıyor.
//---Biz değişkenleri block içerisinde değil de global scope da tanımladığımızda programın çalışma süresi boyunca bellek kaplar. Bu yüzden block scope
//içinde tanımlamak programın daha az bellek kaplaması için önemlidir.-- 