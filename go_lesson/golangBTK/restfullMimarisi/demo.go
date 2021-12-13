package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// json to go tool https://mholt.github.io/json-to-go/
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func _get(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func _post(url string) {

	todo := Todo{1, 500, "Alisveris Yapilacak", false}
	jsonTodo, err := json.Marshal(todo)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
func main() {

	_get("https://jsonplaceholder.typicode.com/todos/1")
	_post("https://jsonplaceholder.typicode.com/todos")

}
