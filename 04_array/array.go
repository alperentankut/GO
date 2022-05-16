package main  

import "fmt"

var arr_1 [3]int
var arr_2 = [5]int{1,2,3,4,5}

func main(){
	arr_3 := make([]int, 3)
	
	arr_3[1] = 2

	fmt.Println(arr_1,arr_2,arr_3)
	fmt.Printf("arr_1 len:%d \n" , len(arr_1))
}

//var keywordu ile array tanımlayabiliyoruz.
//yukarda arr_1 değişkeninin yanına yazdığımız [3]
//ifadesi dizinin boyutunu belirtiyor. Bunu yaptığı-
//mızda go bizim için 3 boyutlu bi integer dizisi oluş-
//turuyor. Fakat herhangi bir değer vermiyor. Eğer değer
//vermemiz gerekiyorsa bi altındaki arr_2 değerinde
//yazdığımız gibi yazıyoruz. Eğer fonksiyon içerisinde
//bir array tanımı yapmak istiyorsak make keywordunu
//kullanıyoruz. make ile bir integer tanımlamak istediğimi
//ve 3 değerlik olacağınız belirtiyoruz. Eğer istediğimiz
//bir indeksine değer vermek istiyorsak köşeli parantez
//ile kaçıncı index olduğunu yazıp sonra değerimizi veriyoruz.