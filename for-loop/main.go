package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		println("Numbers:", i)
	}

	counter := 0
	for counter < 10 {
		println("Infinite loop")
		counter++
		if counter == 5 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}

	data := "name"
	for index, value := range data {
		fmt.Printf("index of data is %d and value is %c\n", index, value)
	}
}