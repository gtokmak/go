package main

import "fmt"

func main() {

	// string --> "", numeric --> 0, bool --> false

	var name1 string            //Declaration
	name1 = "Gokhan"            // Assign
	var name2 string = "Gokhan" //İnitialization
	var name3 = "Gokhan"        // tip belirtmeden
	name4 := "Gokhan"           // short declaration

	var (
		// var tum degiskenlerin basina eklenir
		name5 string            // var name1 string
		name6 string = "Gokhan" // var name6 string = "Gokhan"	a
		name7        = "Gokhan" // var name7 = "Gokhan"
	)

	var name8, name9 = "Gokhan", ""
	name9, name10 := "Aylin", "Gokhan" // daha onceden var olanlari yeni degerlerini atar, tanimlanmamis olan degisleri tanimlar ve deger atar.

	fmt.Println(name1, name2, name3, name4, name5, name6, name7, name8, name9, name10)

	// Go camel case isimlendirme kullanilir
	var coinType string
	var custName string
	// kisaltmalar buyuk harfle yazilir.
	var URL string
	var HTTP string
	var IP string

}
