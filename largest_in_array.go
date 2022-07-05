package main

import "fmt"

func main() {

	var num [100]float64
	var temp int
	fmt.Println("Enter the number of elements")
	fmt.Scan(&temp)

	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number ")
		fmt.Scan(&num[i])
	}

	for j := 0; j < temp; j++ {
		if num[0] < num[j] {
			num[0] = num[j]
		}
	}

	fmt.Print("The largest number in an array is ", num[0])

}
