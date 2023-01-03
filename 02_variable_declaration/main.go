package main

import "fmt"

func main() {

	// var name string = "Alperen"

	// fmt.Println("Hello",name)

	// var name string = "Alperen" 

	name := "Alperen" //Burda değişken tipini belirtmeden
	//direk değeri vermek için := ifadesini kullanırız.

	var age int = 26

	var isMarried bool = true

	var firstName , lastName string = "Alperen" , "Tankut"


	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(firstName , lastName)
}

//Biz func içinde değilsek mutlaka bir keyword ve değişken tipini kullanmak zorundayız. := ile kull-
//andığımız değer atama yöntemi sadece func içinde gerçekleşebiliyor.
//!!! func dışındaysak mutlaka keyword kullanmamız gerekli !!!