package main

import "fmt"

func main() {

	// && -> true apabila true && true
	// || -> true apabila salah satu true
	// ! -> true apabila value false
	nilai := 80
	absensi := true

	fmt.Println("Lulus =" ,!(nilai >= 65 && absensi == true))
}
