package main

import "fmt"

func main() {
	nilai := 100

	switch lulus := false; lulus {
	case false : fmt.Println("Semangatt :))")
	case true : fmt.Println("Selamaat lulus!!")
	default:
		fmt.Println("abc")
	}

	switch {
	case nilai >= 65 : fmt.Println("Lulus")
	case nilai < 65 : fmt.Println("Ga lulus :(")
	}

}
