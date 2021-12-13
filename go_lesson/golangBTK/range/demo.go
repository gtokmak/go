package main

import "fmt"

func main() {
	//key-value
	sehirler := make(map[int]string)
	sehirler2 := map[int]string{1: "Adana", 2: "Adiyaman"}
	sehirler[1] = "Adana"
	sehirler[2] = "Adiyaman"
	sehirler[3] = "Afyon"
	sehirler[6] = "Ankara"
	sehirler[20] = "Denizli"
	sehirler[34] = "Istanbul"

	fmt.Println("Sehirler boyutu: ", len(sehirler))
	deger, varMi := sehirler[40]
	fmt.Printf("deger: %v varMi: %v\n", deger, varMi)
	sehirler[40] = "Konya"
	deger, varMi = sehirler[40]
	fmt.Printf("deger: %v varMi: %v\n", deger, varMi)

	for key, value := range sehirler {
		fmt.Println(key, value)
	}

	for key, value := range sehirler2 {
		fmt.Println(key, value)
	}

}
