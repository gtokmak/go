package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Bakiye yetersiz...")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paraniz hazirlaniyor...")
		hesap -= cekilmekIstenen
	}

	fmt.Println("Hesaptaki para: " + fmt.Sprintf("%v", hesap))
	//fmt.Printf("Hesaptaki para: %v", hesap)

}
