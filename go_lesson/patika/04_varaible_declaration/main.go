//Package clause
package main

//Import statement
import "fmt"

//My Code
func main() {

	//varaible declaration
	var name string
	//Assign
	name = "Gokhan"

	var age int
	age = 30

	var isMarried bool
	isMarried = true

	fmt.Println("Hello ", name)
	fmt.Println(age)
	fmt.Println(isMarried)

	var lastname string = "Tokmak"
	city := "Denizli"
	fmt.Printf("lastname:%v city:%v\n", lastname, city)
	fmt.Printf("Varaible Type -> lastname:%T age:%T isMarried:%T\n", lastname, age, isMarried)

}
