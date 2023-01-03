package main

import "fmt"

func main (){
	var age uint32 = 26

	var isMarried bool = true

	var name string = "Alperen"

	var weight float32 = 72.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Printf("%T\n", name)  // burda "%T metodu ve printf fonksiyonuyla değişkenin tipini yazdırıyoruz. \n bir alt satıra geçmek için."
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
}
// go da bir değişken declare ediyorsak kullanmak zorundayız.

