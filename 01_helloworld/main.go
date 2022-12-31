//package clause
//Bizim go da yazdığımız her kod her zaman paketler halindedir. Biz kodları
//paketler halinde yazarak daha modüler hale getiririz. Daha derli toplu
//olur. Bizim burda yazacağımız bütün kodlar main paketi içerisinde olacak

package main

//Import statement
//Bazı durumlarda kendi paketimiz dışında başkalarının paketlerini veya
//go tarafından hazırlanmış hazır paketleri kullanırız.

import "fmt"

//My code

func main(){
	fmt.Println("Hello world")
}

//go da bir program main paketinin içindeki main fonksiyonundan başlar
//bundan dolayı fonksiyonu main ile yazdık.
//build komutunu yazdığımızda ise bize .exe uzantılı bir dosya veriyor..
