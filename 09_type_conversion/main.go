package main

import "fmt"

func main(){
	x := 10
	y:= 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//Type conversion type(value) => int(y)

	fmt.Println(x + int(y)) //Biz type conversion ile yeni bir değer oluşturuyoruz
	//esas değişkenimizin veri tipini değiştirmiyoruz
}