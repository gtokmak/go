package main

import (
	"fmt"
	"goproje/project"
)

func main() {

	products, _ := project.GetAllProducts()
	fmt.Println(products)

	for _, v := range products {
		fmt.Println(v.ProductName)

	}

	product, err := project.AddProduct()
	fmt.Println(product, err)

}
