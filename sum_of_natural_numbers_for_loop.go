package main

import "fmt"

func main() {

	var n, sum int
	fmt.Print("Enter any Positive number ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Print("Sum : ", sum)
}
