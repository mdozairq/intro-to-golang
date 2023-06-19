package main

import (
	"fmt"
	"net/url"
)

const myurl string= "https://dev.flamapp.com:3000/zingcam/order/admin/flamcards?page=0&page_size=10"

func main(){
	fmt.Println("Handle GoLang url and get details")
	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qParams := result.Query()
	fmt.Printf("Type of Query Params are %T \n", qParams)
	fmt.Println(qParams["page"])

	for _, val := range qParams{
		fmt.Println("Params value is",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "zingcam.dev.flamapp.com",
		Path: "/zingcam/order/admin/flamcards",
		RawQuery: "page=0&page_size=10",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}