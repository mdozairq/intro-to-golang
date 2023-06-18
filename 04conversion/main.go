package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	println("Welcome to the Pizza App")
	println("Please Rate our pizza between 1 and 5 :")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	println("Thanks for Rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}


}