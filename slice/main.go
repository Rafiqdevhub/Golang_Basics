package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = append(numbers, 6)
	fmt.Printf("numbers has data type: %T\n", numbers)
	fmt.Println("slice:",numbers)
	fmt.Println("Length:",len(numbers))
	fmt.Println("Capacity:",cap(numbers))
}