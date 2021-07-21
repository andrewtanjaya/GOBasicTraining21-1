package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Andrew"
	names[1] = "Wiseson"
	names[2] = "Tanjaya"
	fmt.Println(names[0])

	 ages := [10]int{
		18,
		19,
		20,
	}

	fmt.Println(len(ages))
}
