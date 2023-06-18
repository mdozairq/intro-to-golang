package main

import "fmt"

func main(){
	fmt.Println("Struct in GoLang")
	//No Inheritance in golang;
	//No super or parent
	Ozair := User{"Ozair", "mdozairq@gmail.com", true, 16}
	fmt.Println(Ozair)
	fmt.Printf("Details of User: %+v\n", Ozair)
	fmt.Printf("Name: %v and Email: %v", Ozair.Name, Ozair.Email)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}