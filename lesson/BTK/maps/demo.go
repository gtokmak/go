package main

import "fmt"

func main() {
	sehirler := make(map[int]string)
	sehirler2 := map[int]string{1: "Adana", 2: "Adiyaman"}

	sehirler[1] = "Adana"
	sehirler[2] = "Adiyaman"
	sehirler[20] = "Denizli"
	sehirler[6] = "Ankara"

	fmt.Printf("sehirler[%v] = %v", 20, sehirler[20])
	fmt.Printf("sehirler2[%v] = %v", 20, sehirler2[1])
}
