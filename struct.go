package main

import "fmt"

type Person struct{
	FirstName, LastName string
	Age int
}

func main() {
	var andrew Person
	andrew.FirstName = "Andrew"
	andrew.LastName = "Tanjaya"
	andrew.Age = 20

	joko := andrew

	joko.Age = 22
	fmt.Println(andrew.FirstName)
	fmt.Println(andrew.Age)
	fmt.Println(joko.Age)
}
