package main

import (
	"fmt"
	"os"
)

func main() {
	
	// file, error := os.Create("file.txt")
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// defer file.Close()

	// content:= "Hello, World!"
	 // _, errors:=io.WriteString(file, content+"\n")
	// byte, errors:=io.WriteString(file, content+"\n")
	// fmt.Println("Byte Written:",byte)

	// if errors != nil {
	// 	fmt.Println(errors)
	// 	return
	// }

	// fmt.Println("Hello, World!1")
	
	// file, error := os.Open("file.txt")
	// if error != nil {
	// 	fmt.Println("error while opening the file:",error)
	// 	return
	// }

	// defer file.Close()

	// buffer:= make([]byte, 1024)

	// for {
	// 	n, error:=file.Read(buffer)
	// 	if error == io.EOF {
	// 		break
	// 	}
	// 	if error != nil {
	// 		fmt.Println("error while reading the file:",error)
	// 		return
	// 	}
	// 	fmt.Println(string(buffer[:n]))

	// }

	constant, error:=os.ReadFile("file.txt")
	if error != nil {
		fmt.Println("error while reading the file:",error)
		return
	}
	fmt.Println(string(constant))
}