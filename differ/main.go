package main

import "fmt"

func main() {
	fmt.Println("Hello, World!1")
	defer fmt.Println("Hello, World!3")
	defer fmt.Println("Hello, World!...")
	fmt.Println("Hello, World!2")
}