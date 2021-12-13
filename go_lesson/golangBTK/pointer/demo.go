package main

import "fmt"

func demo(sayi *int) {
	*sayi = *sayi + 1
}

func demo1(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayi:", sayilar[0])
}

func main() {
	sayi := 20
	demo(&sayi)
	fmt.Println(sayi)

	sayilar := []int{1, 2, 3}
	demo1(sayilar)
	fmt.Println("Maindeki sayi:", sayilar[0])
}
