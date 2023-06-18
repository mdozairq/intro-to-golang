package main

import "fmt"

var LoginToken int = 99994

func main() {
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("variable is of type smallInt: %T \n", smallInt)

	var smallFloat float32 = 255.94259234823
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type smallFloat: %T \n", smallFloat)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("variable is of type defaultInt: %T \n", defaultInt)

	//implicit type
	var potfolio = "https://ozair.netlify.app"
	fmt.Println(potfolio)
	fmt.Printf("variable is of type potfolio: %T \n", potfolio)

	//no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)
	fmt.Printf("variable is of type numberOfUser: %T \n", numberOfUser)
	numberOfUser=78.7

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type LoginToken: %T \n", LoginToken)

}
