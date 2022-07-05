package main

import (
	"fmt"
	"math"
)

func main() {

	var num [10]float64
	var sum, mean, standardDeviation float64

	fmt.Println("******  Enter 10 elements  *******")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter the %d element : ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}

	mean = sum / 10

	for j := 0; j < 10; j++ {
		standardDeviation += math.Pow(num[j]-mean, 2)
	}

	standardDeviation = math.Sqrt(standardDeviation / 10)

	fmt.Println("The standard Deviation is : ", standardDeviation)

}
