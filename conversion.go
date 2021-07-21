package main

import "fmt"

func main() {
	var angka16 int16 = 128
	angka32 := int8(angka16)
	// -128 -> 127
	fmt.Println(angka32)

	fmt.Println(string("Andrew"[0]))
}
