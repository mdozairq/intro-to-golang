package main

import "fmt"

func main()  {
	fmt.Println("Loops in GoLang")

	days := []string{"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday"}

	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i]);
	// }

	for idx, day := range days{
		fmt.Printf("Index is %v and value is %v \n", idx, day)
	}
}