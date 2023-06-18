package main

import "fmt"

func main() {
	defer fmt.Println("Defers in GoLang 2")
	defer fmt.Println("Defers in GoLang 1")

	fmt.Println("Hello There")
	defer fmt.Println("Defers in GoLang")
	myDefer()
}

func myDefer(){
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}