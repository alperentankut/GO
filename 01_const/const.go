package main

const sabit_1 = "değer 1"

const (
	sabit_2 = "değer 2"
	sabit_3 = "değer 3"
	sabit_4 = "değer 4"
)

const (
	sabit_5 = iota
	sabit_6 
	sabit_7
)

func main(){
	println(sabit_1,sabit_2,sabit_3,sabit_4,sabit_5,sabit_6,sabit_7)
}

//yukardaki iota keywordü bize değerlerimizi sıralı
//şekilde vereceksek bunu otomatik yapmamızı sağlıyor
//yukardaki çıktıda sabit 5 , 6 ve 7 sırasıyla 0,1,2
//şeklinde değer alacaktır.