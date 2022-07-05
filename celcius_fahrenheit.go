package main

import "fmt"

func main() {
	var ftemp float64 = 0
	var ctemp float64 = 0

	fmt.Printf("Enter temperature in Celsius: ")
	fmt.Scanf("%f", &ctemp)
	ftemp = (ctemp * 1.8) + 32
	fmt.Printf("Temperature in Fahrenheit: %.2f", ftemp)

	fmt.Printf("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%f", &ftemp)
	ctemp = (ftemp - 32) / 1.8
	fmt.Printf("Temperature in Celsius: %.2f", ctemp)
}
