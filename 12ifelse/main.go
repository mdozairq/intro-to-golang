package main

import "fmt"

func main(){
	fmt.Println("If Else in GoLang")
	var result string
	loginCount := 10

	if loginCount < 10 {
		result = "The login Count is less than 10"
	} else if loginCount > 10 {
		result = "The login Count in greater than 10"
	} else {
		result = "The login Count is exactly 10"
	}

	fmt.Println(result);

	if num:= 3; num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}