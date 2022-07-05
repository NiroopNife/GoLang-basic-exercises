package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b, res uint32
	fmt.Print("Enter the first number ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number ")
	fmt.Scan(&b)
	res = solveMeFirst(a, b)
	fmt.Print("The result of the two numbers is ", res)
}
