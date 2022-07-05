package main

import "fmt"

func main() {

	var i, j, k, row int

	fmt.Print("Enter Sandglass Number Pattern Rows = ")
	fmt.Scanln(&row)

	fmt.Println("**** Sandglass Number Pattern ****")

	for i = 1; i <= row; i++ {
		for j = 1; j < i; j++ {
			fmt.Printf(" ")
		}
		for k = i; k <= row; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
	for i = row - 1; i >= 1; i-- {
		for j = 1; j < i; j++ {
			fmt.Printf(" ")
		}
		for k = i; k <= row; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
}
