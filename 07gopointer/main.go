package main

import "fmt"

func main() {
	fmt.Println("Welcome to GOlang Pointers")

	var ptr *int
	fmt.Println("Value of Pointer:", ptr)

	myNumber := 32

	var newPtr = &myNumber

	fmt.Println("Value of &myNumber:", myNumber)
	fmt.Println("Value of New Pointer:", newPtr)
	fmt.Println("Value of Actual New Pointer:", *newPtr)

	*newPtr = *newPtr * 2
	fmt.Println("New Value is: ", myNumber)
}
