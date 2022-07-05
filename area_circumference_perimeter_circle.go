package main

import "fmt"

func main() {

	var radius float64
	var area float64
	var perimeter float64 = 0.0
	var circumference float64
	const PI float64 = 3.14

	fmt.Println("Enter radius of Circle : ")
	fmt.Scan(&radius)

	area = PI * radius * radius
	fmt.Println("Area of the circle is ", area)

	circumference = 2 * PI * radius
	fmt.Println("Circumference of the circle is", circumference)

	perimeter = 2 * 3.14 * radius
	fmt.Printf("Perimeter of Circle: %f", perimeter)

}
