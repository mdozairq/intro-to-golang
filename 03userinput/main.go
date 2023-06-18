package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User can Input here"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name of User:")

	//comma ok syntax OR error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thankyou for your Input", input)
	fmt.Printf("Given Input Type are %T \n", input)
}
