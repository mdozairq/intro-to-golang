package main

import "fmt"

func main(){
	fmt.Println("Struct in GoLang")
	//No Inheritance in golang;
	//No super or parent
	Ozair := User{"Ozair", "mdozairq@gmail.com", true, 16}
	fmt.Println(Ozair)
	fmt.Printf("Details of User: %+v\n", Ozair)
	Ozair.GetStatus();
	Ozair.UpdateUserEmail();
	fmt.Printf("Details of User: %+v\n", Ozair)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is User Active", u.Status)
}

func (u User) UpdateUserEmail(){
	u.Email="ozair@dev.in"
	fmt.Println("New Email is:", u.Email)
}