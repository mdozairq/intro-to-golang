package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to Files in GoLang")
	content := "I am learning working with files in GoLang"
	file, err := os.Create("./work.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	fmt.Println("length of character in txt file:",length)
	defer file.Close()

	readFile("./work.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text inside the file are is \n", string(databyte))
}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}