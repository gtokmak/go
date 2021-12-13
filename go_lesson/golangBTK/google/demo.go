package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRequest(url string) {

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	//fmt.Println(response.Header)

	for k, v := range response.Header {
		fmt.Println(k, v)
	}

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}

func main() {
	getRequest("https://www.google.com.tr")
}
