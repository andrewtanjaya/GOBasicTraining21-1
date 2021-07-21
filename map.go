package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama" : "Andrew",
		"alamat" : "Jalan jeruk",
	}
	orang["telp"] = "081321312312"

	fmt.Println(orang["telp"])
	fmt.Println(len(orang))
	delete(orang, "telp")
	fmt.Println(len(orang))

}
