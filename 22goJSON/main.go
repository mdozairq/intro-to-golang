package main

import (
	"encoding/json"
	"fmt"
)

type post struct {
	UserId int      `json:"userId"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Tags   []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Making JSON in GOlang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	userPosts := []post{
		{2, "Post1 Title", "PostBody on Post1", []string{"Happy", "New", "Tech"}},
		{3, "Post2 Title", "PostBody on Post2", []string{"GoLang", "Stuff", "Fun"}},
		{1, "Post3 Title", "PostBody on Post3", nil},
	}

	//package userPosts to JSON
	// finalJson, err := json.Marshal(userPosts)
	finalJson, err := json.MarshalIndent(userPosts, "", "\t")
	checkNilError(err)

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(
		`{
                "userId": 3,
                "title": "Post2 Title",
                "body": "PostBody on Post2",
                "tags": [
                        "GoLang",
                        "Stuff",
                        "Fun"
                ]
        }`)

	var postData post

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &postData)
		fmt.Printf("%#v\n", postData)
	} else {
		fmt.Println("JSON is not Valid")
	}

	//some case

	var myPostData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myPostData)
	fmt.Printf("%#v\n", myPostData)

	for K, V := range myPostData {
		fmt.Printf("Key is: %v and Value is: %v and also its type is: %T \n", K, V, V )
	}
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
