package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to GOlang Slice")
	var fruitList = []string{"APPLE", "TOMATO", "PEACH", "GUAVA"}
	fmt.Printf("Type of fruitList: %T \n", fruitList)
	fruitList = append(fruitList, "GRAPES", "BANANA")
	fmt.Println("Array vegList are: ", fruitList[2:5])

	highScore := make([]int, 4)
	highScore[0] = 239
	highScore[1] = 235
	highScore[2] = 236
	highScore[3] = 237

	fmt.Println(highScore)
	// sort.Ints(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"REACT","NODE", "NEST",  "NEXT", "EXPRESS"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	
}