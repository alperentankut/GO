package main

import "fmt"

func main(){

	// var studentName string = "John Doe"
	// var grade int = 77
	// var isPassed bool = true

	// var studentName , grade , isPassed = "John Doe" , 77 , true

	studentName , grade , isPassed := "John Doe" , 77 , true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

	//GO da bir değişken declare edildikten sonra tekrar declare edilemez. Python da bu böyle değildir declare ettiğimiz bir değişkeni tekrar declare
	//edebiliriz. Bu da GO yı statically type bir dil yapar. Fakat python dynamically type bir dildir.

	// var studentName string = "Alperen"
	// var studentName string = "Mehmet"
	//--- Yukarda iki kere declare etmeye çalıştığımız için hata alırız. Fakat ikinci değişkeni studentName = "Mehmet" şeklinde yazsaydık hata almazdık 
	//çünkü yeni bir değer atamış oluyoruz.--- 
}