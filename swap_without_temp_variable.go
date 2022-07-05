package main

import "fmt"

func main() {

	var first, second int

	fmt.Print("Enter the first number ")
	fmt.Scan(&first)
	fmt.Print("Enter the second number ")
	fmt.Scan(&second)

	first = first - second
	second = first + second
	first = second - first

	fmt.Printf("First number is %d, Second number is %d", first, second)
}
