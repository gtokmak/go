package main

import (
	"fmt"
	"time"
)

func tekSayilar() {

	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek sayı:", i)
		time.Sleep(1 * time.Second)
	}
}

func çiftSayilar() {

	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift sayı:", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go tekSayilar()
	go çiftSayilar()
	time.Sleep(5 * time.Second)

}
