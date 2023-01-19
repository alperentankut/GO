package main

import "fmt"

func main(){
// 	x := 10
// 	y:= 10.0

// 	fmt.Printf("%v %T\n", x, x)
// 	fmt.Printf("%v %T\n", y, y)

// 	//Type conversion type(value) => int(y)

// 	fmt.Println(x + int(y)) //Biz type conversion ile yeni bir değer oluşturuyoruz
// 	//esas değişkenimizin veri tipini değiştirmiyoruz
//  int veri tipleri arasında bile tip uyuşmazlığı hatası alabiliyoruz. int8 ve
//  int16 iki veri tipini toplamaya kalktığımızda hata alırız. 
  
	// x := 10
	// y := 10.0
	// fmt.Printf("%v %T\n", x, x)
 	// fmt.Printf("%v %T\n", y, y)

	// fmt.Println(float64(x)+y)

	var x int8 = 127
	var y int16

	y = int16(x)

	fmt.Println(y)

	//yaptığımız işlemin senaryosuna göre float64 veya int değerlerinden
	//hangisine çevireceğimize karar veriyoruz. Mantık olarak daha çok veri
	//kapsayan değere dönüştürmek daha avantajlı olabilir.

	// --- String bir ifadeyi type conversion ile integer yapamayız ---
}