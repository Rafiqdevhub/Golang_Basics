package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Salam, Enter your name: ")

	// var name string

	// fmt.Scanln(&name)

	// fmt.Println("Hello", name)

	reader:=bufio.NewReader(os.Stdin)
	name, _:=reader.ReadString('\n')
	fmt.Println("Hello", name)
}