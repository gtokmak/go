package funtions

import "fmt"

func SelamVer(adi string) {
	fmt.Println("Merhaba", adi)
}

func Toplama(sayi1 int, sayi2 int) int {

	return sayi1 + sayi2
}

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {

	return (sayi1 + sayi2), (sayi1 - sayi2), (sayi1 * sayi2), float32(sayi1 / sayi2)
}

func Topla(sayilar ...int) int {

	toplam := 0

	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}

	return toplam
}
