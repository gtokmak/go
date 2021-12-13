package slice

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	isimler[0] = "Gokhan"
	isimler[1] = "Aylin"
	isimler[2] = "mustafa"

	fmt.Println(isimler)
	fmt.Println(len(isimler))

	isimler = append(isimler, "behiye")

	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
