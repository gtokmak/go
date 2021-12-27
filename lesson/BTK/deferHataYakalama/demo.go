package main

import (
	"fmt"
	"os"
)

func A() {

	fmt.Println("a fonksiyonum calisti")
}

func C() {

	fmt.Println("c fonksiyonum calisti")
}

func B() {
	defer A() // defer b fonkisyonun bittikten sonra calisir ve stack yapisinda saklanir
	defer C() // defer b fonkisyonun bittikten sonra calisir ve stack yapisinda saklanir
	// stack en son c olacagi icin ilk C calisir sonrasinda A calisir
	// b fonksiyonunda hata da olsa c ile a calisir.
	fmt.Println("b fonksiyonum calisti")

}

func main() {
	B()

	f, err := os.Open("demo.txt")

	if err != nil {
		
		fmt.Println("hata:", err)
	} else {

		fmt.Println("dosya acildi", f.Name())

	}
}
