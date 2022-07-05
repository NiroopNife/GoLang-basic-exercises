package main

import "fmt"

func main() {

	fmt.Println("--------------------------")
	fmt.Println("Enter an integer")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is an odd number")
	}

	fmt.Println("--------------------------")
	var odnum int

	fmt.Print("Enter the Number to Print Odd's = ")
	fmt.Scanln(&odnum)

	for x := 1; x <= odnum; x++ {
		if x%2 != 0 {
			fmt.Print(x, "\t")
		}
	}

	fmt.Println("--------------------------")
	var evnum int

	fmt.Print("Enter the Number to Print Even's = ")
	fmt.Scanln(&evnum)

	fmt.Println("Even Numbers from 1 to ", evnum, " are = ")
	for i := 1; i <= evnum; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		}
	}

	fmt.Println("--------------------------")
	var eonum, eventotal, oddtotal int

	fmt.Print("Enter the Number to find Even and Odd Sum = ")
	fmt.Scanln(&eonum)

	eventotal = 0
	oddtotal = 0

	for x := 1; x <= eonum; x++ {
		if x%2 == 0 {
			eventotal = eventotal + x
		} else {
			oddtotal = oddtotal + x
		}
	}
	fmt.Println("\nSum of Even Numbers from 1 to ", eonum, " = ", eventotal)
	fmt.Println("\nSum of Odd Numbers from 1 to ", eonum, "  = ", oddtotal)

}
