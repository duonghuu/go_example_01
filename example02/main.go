package main

import "fmt"

func main() {
	// iết:

	// If kiểm tra số chẵn/lẻ

	// Vòng for từ 1 đến 10

	// For-range duyệt slice
	// var myNumber = 10
	// if myNumber%2 == 0 {
	// 	fmt.Println("number is even")
	// } else {
	// 	fmt.Println("number is odd")

	// }

	// index := 1
	// for {
	// 	fmt.Printf("number: %d\n", index)
	// 	if index >= 10 {
	// 		break
	// 	}

	// 	index++

	// }
	// for i := 1; i < 11; i++ {
	// 	fmt.Printf("number: %d\n", i)
	// }

	mySlice := []int{0, 2, 3, 4, 9, 0, 3, 1, 2, 9}
	for _, value := range mySlice {
		fmt.Printf("number: %d\n", value)
	}
	fmt.Printf("len: %d\n", len(mySlice))
}
