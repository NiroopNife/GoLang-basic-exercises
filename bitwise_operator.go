package main

import "fmt"

func main() {
	var num int = 0
	var num1 int = 0
	var num2 int = 0
	var leftShift int = 0
	var rightShift int = 0
	var bitwiseAnd int = 0
	var bitwiseOr int = 0
	var bitwiseXor int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("Enter number1: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter number2: ")
	fmt.Scanf("%d", &num2)

	leftShift = num << 3
	fmt.Printf("Left shift Bitwise : %d", leftShift)

	rightShift = num >> 3
	fmt.Printf("Right shift Bitwise : %d", rightShift)

	bitwiseAnd = num1 & num2
	fmt.Printf("Result: %d", bitwiseAnd)

	bitwiseOr = num1 | num2
	fmt.Printf("Result: %d", bitwiseOr)

	bitwiseXor = num1 ^ num2
	fmt.Printf("Result: %d", bitwiseXor)

}
