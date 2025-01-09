package main

import "fmt"

func sayHello() {
	fmt.Println("I am saying hello from the function")
}

func add(a, b int) (result int) {
	result = a + b
	return 
}

func main() {
	sayHello()

	result := add(10, 20)
	fmt.Println("Result is: ", result)
}