package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Handling Web Request in GoLang")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	// url := "https://zingcam.dev.flamapp.com/zingcam/order/admin/flamcards?page=0&page_size=10"
	url := "https://jsonplaceholder.typicode.com/posts/1"
	response, err := http.Get(url)

	checkNilError(err)
	
	defer response.Body.Close()
	fmt.Println("Response Status:", response.StatusCode)
	fmt.Println("Response content:", response.ContentLength)

	var responseString strings.Builder
	data, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(data)
	// content := string(data)
	// fmt.Println(content)
	fmt.Println("ContentLength: ", byteCount)
	fmt.Println("content:", responseString.String())
}

func PerformPostJSONRequest(){
	url := "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
	{
		"title": "foo",
		"body": "bar",
		"userId": 1
	  }
	`)

	response, err := http.Post(url,"application/json", requestBody)

	checkNilError(err)

	defer response.Body.Close()
	fmt.Println("Response Status:", response.StatusCode)
	var responseString strings.Builder
	data, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(data)
	fmt.Println("ContentLength: ", byteCount)
	fmt.Println("content:", responseString.String())
	
}

func PerformPostFormRequest(){
	myUrl := "https://jsonplaceholder.typicode.com/posts"

	formData := url.Values{}
	formData.Add("title","foo")
	formData.Add("body","bar")
	formData.Add("userId", "23")



	response, err := http.PostForm(myUrl, formData)

	checkNilError(err)

	defer response.Body.Close()
	fmt.Println("Response Status:", response.StatusCode)
	var responseString strings.Builder
	data, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(data)
	fmt.Println("ContentLength: ", byteCount)
	fmt.Println("content:", responseString.String())
	
}



func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}