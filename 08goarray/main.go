package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Array in golang")

	var arrayItems [4]string
	arrayItems[0] = "Zero"
	arrayItems[1] = "One"
	arrayItems[3] = "Three"

	fmt.Println("Array List are: ", arrayItems)
	fmt.Println("Length of Array List are: ", len(arrayItems))

	var vegList = [3]string{"Potato","Brinjal", "Beans"}

	fmt.Println("Array vegList are: ", vegList)
	fmt.Println("Length of vegList List are: ", len(vegList))
}