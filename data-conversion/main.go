package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num = 10

	// fmt.Println("Value of num:", num)
	// fmt.Printf("Type of num is %T\n", num)

	// var data float64=float64(num)
	// data=data+5.5
	// fmt.Println("Value of data:", data)
	// fmt.Printf("Type of data is %T\n", data)

	// num = 123
	// str:= strconv.Itoa(num)
	// fmt.Println("Value of str:", str)
	// fmt.Printf("Type of str is %T\n", str)

	// name := "123"
	// num, _ = strconv.Atoi(name)
	// num = num + 5
	// fmt.Println("Value of num:", num)
	// fmt.Printf("Type of num is %T\n", num)

	num_str := "12.3"
	number_float, _ := strconv.ParseFloat(num_str, 64)
	fmt.Println("Value of number_float:", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}