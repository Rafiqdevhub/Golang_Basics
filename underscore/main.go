package main

import "fmt"

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }


func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "cannot divide by zero"
	}
	return a / b, "nil"
}

func main() {
	println("Hello, World!")
	fmt.Println("The idea is to have a single import path for all packages")
	ans, _:= divide(10, 0)
	fmt.Println("the answer is ", ans)
}