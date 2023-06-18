package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.May, 16, 21, 24, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println("01-02-2006 15:04:05 Monday")
}
