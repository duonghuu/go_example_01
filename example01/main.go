package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sayHello() {

}

func main() {
	var name string = "eric"
	var age int = 10
	var gender bool = false
	const PI = 3.14
	genderText := ""
	if gender == false {
		genderText = "Nam"
	} else {
		genderText = "Nữ"
	}

	fmt.Printf("%s %d %s\n", name, age, genderText)
	fmt.Printf("answer: %.2f", PI)
}
