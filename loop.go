package main

import "fmt"

func main() {
	// init statement, post statement
	//counter := 1


	numbers := [...]int{10,40,232}

	for counter:= 0; counter < len(numbers) ; counter++{
		fmt.Println(numbers[counter])
	}

	for index, number := range numbers{
		if number == 40 {
			continue
		}
		fmt.Println("Index ", index , "=" , number)

	}

}
