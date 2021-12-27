package main

import (
	"fmt"
	s "strings"
)

func main() {

	isim := "Gokhan kkkkk"
	fmt.Println(s.Count(isim, "k"))    // kac tane oldugunu bulur
	fmt.Println(s.Contains(isim, "o")) // içeriyorsa true doner
	fmt.Println(s.Index(isim, "kh"))   // ilk basladigi index döner yoksa -1
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
	fmt.Println(isim)
	fmt.Println(s.HasPrefix(isim, "Gok")) //  ile basliyor mu?
	fmt.Println(s.HasSuffix(isim, "an"))  //	ile bitiyor mu?

	soyisim := []string{"t", "o", "k", "m", "a", "k"}
	sonuc := s.Join(soyisim, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3)) // -1 varsa hepsini 3 yaparsak sadece üc tanesini degistir
	fmt.Println(s.Split(sonuc, "*"))
}
