package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)

	formattedTime := currentTime.Format("2006-01-02")
	fmt.Println(formattedTime)

	formattedTime = currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime)
}