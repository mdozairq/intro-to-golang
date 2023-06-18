package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Switch and Case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1
	fmt.Println("Value of Dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead. \n", diceNumber, diceNumber)
	case 2:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead. \n", diceNumber, diceNumber)
	case 3:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead. \n", diceNumber, diceNumber)
		fallthrough
	case 4:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead. \n", diceNumber, diceNumber)
		fallthrough
	case 5:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead. \n", diceNumber, diceNumber)
	case 6:
		fmt.Printf("Dice Value: %v, hence you can move %v step ahead or Start game and also repeat. \n", diceNumber, diceNumber)
	default:
		fmt.Printf("Not defined!")
	}
	
}