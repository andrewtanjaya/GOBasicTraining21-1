package main

import "fmt"

func sayHello(){
	fmt.Println("Hello Andrew👋")
}

func sayHelloToSomeone(name string , hello func()){
	hello()
	fmt.Println("Hello", name)
}

func getHello(name string) string{
	return "Hello " + name
}

// named return value
func getData() (nama string, umur int8){
	nama = "Andrew"
	umur = 20
	return
}

// variadic function
func sumAll(numbers ...int) int{
	total := 0
	for _, number := range numbers{
		total+=number
	}

	return total
}

func main() {
	//numbers := [...]int{10,20,100,1}
	//slice1 := numbers[:]
	//fmt.Println(sumAll(slice1...))

	// function as value
	sayHelloToSomeone("Andrew", sayHello)
}
