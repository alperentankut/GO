package main

import "fmt"


var map_1 map[int]string

func main() {
	//map = make(map[int]string)

	map_2 := make(map[int]int)
	map_3 := make(map[int]int, 3)

	fmt.Println(map_1 , map_2 , map_3)
	
}


// mapler sıra garantisi olmadan memoryde adres ve 
//referans mantalitesiyle çalışıyor. mapten sonra key
//tipini ve value tipiyle bir map tanımı yapabiliyoruz. 