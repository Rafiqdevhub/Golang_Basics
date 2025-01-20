package main

import (
	"fmt"
	"strings"
)

func main() {
	// data := "Ali, Ahmed , 123, 456, 789"
	// parts := strings.Split(data, ",")
	// fmt.Printf("parts: %v\n", parts)

	// dataFormat:= "Ali.Ahmed.123.45.789"
	// dataParts:= strings.Split(dataFormat, ".")
	// fmt.Printf("dataParts: %v\n", dataParts)
	
	// countsString:="one, two, three, four, five, five"
	// count:=strings.Count(countsString, "five")
	// fmt.Printf("count: %v\n", count)

	// trimpingString:="   Ali Ahmed   "
	// trimString:=strings.TrimSpace(trimpingString)
	// fmt.Printf("trimString: %v\n", trimString)

	firstName:="Ali"
	lastName:="Ahmed"
	// fullName:=strings.Join([]string{firstName, lastName}, " ")
	fullName := strings.Join([]string{firstName,"Khan", lastName}, " ")
	fmt.Printf("fullName: %v\n", fullName)
}