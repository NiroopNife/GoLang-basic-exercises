package main

import "fmt"

var val1 int = 0
var val2 int = 0

func setVal1() {
	val1 = 10
}

func setVal2() {
	val1 = 20
}

func sum() {
	fmt.Println("Sum: ", val1+val2)
}

func main() {
	defer sum()
	defer setVal2()
	setVal1()
}
