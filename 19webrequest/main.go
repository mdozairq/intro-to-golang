package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Handling web request using GoLang")
	url := "https://lco.dev"
	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Type of Response: %T \n", response)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	content := string(data)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
