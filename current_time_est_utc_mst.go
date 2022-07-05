package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	z, _ := t.Zone()
	fmt.Println("Zone : ", z, "Time : ", t)

	location, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Zone : ", location, "Time : ", t.In(location))

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ZONE : ", loc, " Time : ", now)

	loc, _ = time.LoadLocation("MST")
	now = time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now)
}
