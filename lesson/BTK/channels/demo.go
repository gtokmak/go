package main

import (
	"fmt"
)

func tekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek sayı:", i)
		toplam += i
	}
	tekSayiCn <- toplam
}

func çiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift sayı:", i)
		toplam += i
	}
	ciftSayiCn <- toplam
}

func main() {
	tekCn, ciftCn := make(chan int), make(chan int)

	go tekSayilar(tekCn)
	go çiftSayilar(ciftCn)

	tekSayiToplam, ciftSayiToplam := <-tekCn, <-ciftCn
	fmt.Println("Carpim:", tekSayiToplam*ciftSayiToplam)

}
