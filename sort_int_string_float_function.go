package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{50, 90, 30, 10, 50}
	fmt.Println(num)
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num)
	}
	fmt.Println(num)
	fmt.Println(sort.SearchInts(num, 30))

	text := []string{"Nife", "Niroop", "Arsenal", "RCB"}
	fmt.Println(text)
	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	fmt.Println(text)
	fmt.Println(sort.SearchStrings(text, "Arsenal"))

	val := []float64{2.4, 4.5, 3.45, 3.2, 2.35, 5.6, 5.7}
	fmt.Println(val)
	if sort.Float64sAreSorted(val) == false {
		sort.Float64s(val)
	}
	fmt.Println(val)
	fmt.Println(sort.SearchFloat64s(val, 4.5))

}
