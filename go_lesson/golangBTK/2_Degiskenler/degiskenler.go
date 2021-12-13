package main

import "fmt"

func main() {
	var metin string = "Merhaba Türkiye..."
	var oran int = 50
	var fvar float32 = 7.7

	fmt.Println(metin)
	fmt.Println(oran)
	fmt.Println(fvar)

	kdv := 25
	fmt.Printf("veri tipi: %T\n", kdv)
	kdv1 := 25.5
	fmt.Printf("veri tipi: %T\n", kdv1)

	var durum bool = false

	fmt.Println(durum)
	fmt.Println(!durum)

}
