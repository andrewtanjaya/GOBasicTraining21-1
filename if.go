package main

import "fmt"

func main() {

	if umur := 1; umur <= 1 {
		fmt.Println("Masih Bayi")
	}else if umur <= 5{
		fmt.Println("Balita")
	}else if umur <= 17 {
		fmt.Println("Remaja")
	}else {
		fmt.Println("Dewasa")
	}

}
