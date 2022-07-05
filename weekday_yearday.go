package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	fmt.Println(t.YearDay())
	fmt.Println(t.Weekday())
}
