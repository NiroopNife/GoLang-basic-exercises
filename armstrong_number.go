package main

import "fmt"

func main() {

	var number, remainder, tempNumber int
	var result int = 0

	fmt.Println("Enter any integer")
	fmt.Scanln(&number)

	tempNumber = number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10

		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Println(number, "is an Armstrong number")
	} else {
		fmt.Println(number, "is not an Armstrong number")
	}

}
