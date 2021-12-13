package main

import "fmt"

type product struct {
	name       string
	unitePrice float64
	brand      string
}

type costumer struct {
	firstName string
	lastName  string
	age       int
}

func (c costumer) save() {
	fmt.Println("eklendi : ", c.firstName)
}

func (c costumer) update(firstName string) {
	c.firstName = firstName
	fmt.Println("güncellendi : ", c.firstName)
}

func main() {

	kalem := product{"tukenmez", 500, "faber"}
	silgi := product{name: "tukenmez", brand: "faber"} // deger girilmeyen alanlar varsayşlan deger atanir
	fmt.Println(kalem)
	fmt.Println(silgi)

	c := costumer{firstName: "gokhan", lastName: "tokmak", age: 30}
	c.save()
	c.update("Aylin")

}
