package main

import "fmt"

func main() {

	var num [4]int
	var temp, sum, avg int

	fmt.Print("Enter the number of elements : ")
	fmt.Scan(&temp)

	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number")
		fmt.Scan(&num[i])

		sum += num[i]
	}

	avg = sum / temp
	fmt.Printf("The average of entered %d elements is %d", temp, avg)

}
