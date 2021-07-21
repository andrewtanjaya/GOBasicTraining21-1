package test

import "fmt"

// diawali dengan huruf besar -> boleh dipake di package lain
// diawali dengan huruf kecil -> ga boleh dipake di package lain
func abc(){
	fmt.Println("abc")
}

func Def(){
	fmt.Println("def")
}
