package main

import "fmt"

func main(){
	fmt.Println("Welcome to function in golang")
	greeter()
	fmt.Println(adder(3,4))
	result, name := proAdder(2,3,4,5,54)
	fmt.Println("Total of given number",result, name)
}

func greeter(){
	fmt.Println("In the function Greeter")
	// greeter()
	// return "Hye to Function in GoLang"
}

func adder(a int,b int) int{
	fmt.Println("Adder")

	return a+b
}

func proAdder(values ...int) (int, string) {
	total:=0;
	for _, val := range values{
		total+=val
	}

	return total, "ozair"
}

// func () {
// 	fmt.println("Anonymous Func")
// }()