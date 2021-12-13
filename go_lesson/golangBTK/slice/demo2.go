package slice

import "fmt"

func Demo2() {

	sehirler := []string{"Ankara", "İzmir", "İstanbul"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

}
