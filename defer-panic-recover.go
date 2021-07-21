package main

import "fmt"

func shutdownApp (){
	message := recover()
	if message != nil{
		fmt.Println("Error Message : ", message)
	}
	fmt.Println("Application Shutdown")
}

func runApplication(value int) {
	defer shutdownApp()
	fmt.Println("Application Running")
	result := 10/value
	fmt.Println(result)
}

func testPanic(error bool){
	defer shutdownApp()
	if error {
		panic("Application Error")
	}
}

func main() {
	fmt.Println("Application Started")
	//runApplication(0)
	testPanic(true)
	fmt.Println("App Works")
}
