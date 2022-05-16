package main  

import "fmt"

var arr_1 [3]int
var arr_2 = [5]int{1,2,3,4,5}
var slc_1 []int

func main(){

	slc_2 := make ([]int , 0 ,3)

	// slc_2[0] = 2 --Hata mesajı alırız--

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)

	arr_3 := make([]int, 3)
	
	arr_3[1] = 2

	fmt.Println(arr_1,arr_2,arr_3)
	fmt.Printf("arr_1 len:%d \n" , len(arr_1))
	fmt.Println(slc_1 , slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n" , len(slc_1),cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n" , len(slc_2),cap(slc_2))
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

// --Slices--

//go da bir de slice dediğimiz yapılar vardır. Bunlarda aynı
//arrayler gibidir fakat arraylerden ayıran tek fark arraylerde
//boyut bellidir fakat slice lar genişleyebilir yapılardır.
//yazdığımız cap keywordu ise bir slices ın ne kadar kapasitesi
//olduğunu öğreniyoruz. Slice yazarken 3 parametre giriyoruz. 
//1. parametre veri tipi 2. parametre dizinin boyutunu
//3. parametre ise arrayın kapasitesini belirtiyor. Biz slc_2 
//slices ımıza slc_2[0]=2 şeklinde bişey yazıp 0. elemanı 2 yapmak
//istersek hata alırız çünkü bu dizimizin boyutunu 0 olarak belirttik.
//buna bi eleman eklemek istersek append keywordunu kullanıyoruz. 
//verdiğimiz kapasiteyi aşacak şekilde bi ekleme yaparsak kapasite 2 kat
//artarak artmaya devam eder. Eğer cap değeri vermezsek default olarak 2
//değerini alır ve 2 nin katları şeklinde artarak devam eder.