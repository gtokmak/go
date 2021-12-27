package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Pruduct struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryID  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Pruduct, error) {

	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Pruduct
	json.Unmarshal(bodyBytes, &products)
	return products, nil

}

func AddProduct() (Pruduct, error) {
	product := Pruduct{ProductName: "Microphone-3", CategoryID: 1, UnitPrice: 7777}
	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return Pruduct{}, err
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Pruduct
	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil
}
