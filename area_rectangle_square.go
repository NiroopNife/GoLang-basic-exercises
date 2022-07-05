package main

import "fmt"

func main() {

	var area, l, b int
	fmt.Print("Enter the length of the Rectangle ")
	fmt.Scan(&l)
	fmt.Print("Enter the breadth of the Rectangle ")
	fmt.Scan(&b)

	area = l * b
	fmt.Println("The area of Rectangle is ", area)

	fmt.Print("Enter the length of the Square ")
	fmt.Scan(&l)

	area = l * l
	fmt.Println("Area of Square is ", area)
}
