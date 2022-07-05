package main

import "fmt"

func main() {

	var first, second, third int

	fmt.Println("Enter the first number ")
	fmt.Scan(&first)
	fmt.Println("Enter the second number ")
	fmt.Scan(&second)
	fmt.Println("Enter the third number ")
	fmt.Scan(&third)

	if first >= second && first >= third {
		fmt.Print(first, " is the greatest among the three numbers")
	} else if second >= first && second >= third {
		fmt.Print(second, " is the greatest among the three numbers")
	} else if third >= first && third >= second {
		fmt.Print(third, " is the greatest among the three numbers")
	}

}
