package main

import "fmt"

func main() {
	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"Agust",
		"Sept",
		"Oct",
		"Nov",
		"Des",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(months)
	slice1[0] = "May_Updated"

	fmt.Println(months)
	fmt.Println(slice1)
}
