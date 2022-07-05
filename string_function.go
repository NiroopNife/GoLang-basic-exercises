package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println("--------------------------")
	fmt.Println(strings.Compare("India", "Indiana"))
	fmt.Println(strings.Compare("Indiana", "India"))
	fmt.Println(strings.Compare("India", "India"))

	fmt.Println("--------------------------")
	fmt.Println("Contains function: ", strings.Contains("India", "ia"))
	fmt.Println("Contains function: ", strings.Contains("India", "ma"))

	fmt.Println("--------------------------")
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))

	fmt.Println("--------------------------")
	fmt.Println("Count Function: ", strings.Count("india", "i"))

	fmt.Println("--------------------------")
	fmt.Println("Equal Fold Function: ", strings.EqualFold("Cat", "cAt"))
	fmt.Println("Equal Fold Function: ", strings.EqualFold("India", "Indiana"))

	fmt.Println("--------------------------")
	testString := "India is a country which is having soo many states"
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}

	fmt.Println("--------------------------")
	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	strArray := strings.FieldsFunc(`India major cities - Bengaluru, Chennai, Hyderabad, Mumbai, Delhi`, x)
	for _, v := range strArray {
		fmt.Println(v)
	}
	fmt.Println("*****Split by number*****")
	y := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	mysuruArray := strings.FieldsFunc(`1 Mysuru Palace.2 Mysuru Zoo.3 Chamundi Hills.4 KRS`, y)
	for _, w := range mysuruArray {
		fmt.Println(w)
	}

	fmt.Println("--------------------------")
	fmt.Println("Has Prefix: ", strings.HasPrefix("India", "In"))
	fmt.Println("Has Prefix: ", strings.HasPrefix("India", "Na"))

	fmt.Println("--------------------------")
	fmt.Println("Has Suffix: ", strings.HasSuffix("India", "ia"))
	fmt.Println("Has Suffix: ", strings.HasSuffix("India", "la"))

	fmt.Println("--------------------------")
	fmt.Println(strings.Index("India", "Ind"))
	fmt.Println(strings.Index("India", "ind"))
	fmt.Println(strings.Index("India", "i"))
	fmt.Println(strings.Index("India", "Jap"))
	fmt.Println(strings.Index("India-124", ""))

	fmt.Println("--------------------------")
	fmt.Println(strings.IndexAny("India", "japan")) // I position
	fmt.Println(strings.IndexAny("India", "dia"))   // d position

	fmt.Println("--------------------------")
	var s, t, u byte
	t = 'd'
	s = 1
	u = '1'

	fmt.Println(strings.IndexByte("India", t))
	fmt.Println(strings.IndexByte("INDIA", t))
	fmt.Println(strings.IndexByte("5221-INDIA", s))
	fmt.Println(strings.IndexByte("5221-INDIA", u))

	fmt.Println("--------------------------")
	var m, n, o rune
	n = 'd'
	m = 1
	o = '1'

	fmt.Println(strings.IndexRune("India", n))
	fmt.Println(strings.IndexRune("INDIA", n))
	fmt.Println(strings.IndexRune("5221-INDIA", m))
	fmt.Println(strings.IndexRune("5221-INDIA", o))

	fmt.Println("--------------------------")
	fmt.Println("Join Function: ", strings.Join([]string{"Karnataka", "Andra Pradesh", "Tamil Nadu"}, "-"))
	fmt.Println("Join Function: ", strings.Join([]string{"Karnataka", "Andra Pradesh", "Tamil Nadu"}, " "))

	fmt.Println("--------------------------")
	fmt.Println(strings.LastIndex("India", "i"))
	fmt.Println(strings.LastIndex("japanese", "a"))
	fmt.Println(strings.LastIndex("1234567890 1234567890", "00"))
	fmt.Println(strings.LastIndex("1234567890 1234567890", "123"))

	fmt.Println("--------------------------")
	fmt.Println("Repeat Function: ", strings.Repeat("India", 3))

	fmt.Println("--------------------------")
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 1))
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 2))
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 3))
	fmt.Println("Replace Function: ", strings.Replace("India", "dia", "idn", 1))

	fmt.Println("--------------------------")
	fmt.Println("Split Function: ", strings.Split("Karnataka-Andra Pradesh-Tamil Nadu", "-"))

	fmt.Println("--------------------------")
	strSlice := strings.SplitN("a,b,c", ",", 0)
	fmt.Println(strSlice)
	strSlice = strings.SplitN("a,b,c", ",", 1)
	fmt.Println(strSlice)

	fmt.Println("--------------------------")
	stringSlice := strings.SplitAfter("a,b,c", ",")
	fmt.Println(stringSlice)
	stringSlice = strings.SplitAfter("abacadaeaf", "a")
	fmt.Println(stringSlice)

	fmt.Println("--------------------------")
	fmt.Println(strings.Title("towns and cities"))
	fmt.Println(strings.Title("Germany is a Western European country with a landscape of forests, rivers, mountain ranges and North Sea beaches."))

	fmt.Println("--------------------------")
	fmt.Println(strings.ToTitle("towns and cities"))
	fmt.Println(strings.ToTitle("Germany is a Western European country with a landscape of forests, rivers, mountain ranges and North Sea beaches."))

	fmt.Println("--------------------------")
	fmt.Println("To Lower Function", strings.ToLower("India"))

	fmt.Println("--------------------------")
	fmt.Println("To Upper Function", strings.ToUpper("India"))

	fmt.Println("--------------------------")
	fmt.Println(strings.Trim("0120 2510", "0"))
	fmt.Println(strings.Trim("abcd axyz", "a"))
	fmt.Println(strings.Trim("abcd axyz", "A"))

	fmt.Println("--------------------------")
	fmt.Println(strings.TrimLeft("0120 2510", "0"))
	fmt.Println(strings.TrimLeft("abcd axyz", "a"))

	fmt.Println("--------------------------")
	fmt.Println(strings.TrimRight("0120 2510", "0"))
	fmt.Println(strings.TrimRight("abcd axyz", "a"))

	fmt.Println("--------------------------")
	fmt.Println(strings.TrimSpace(" New Zealand is a country in the southwestern Pacific Ocean "))
	fmt.Println(strings.TrimSpace(" \t\n  New Zealand is a country in the southwestern Pacific Ocean \t\n "))

	fmt.Println("--------------------------")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter Second String: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first + second)
}
