package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{"Ali", 30, true}
	fmt.Println(person)

	jsonData, error:=json.Marshal(person)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(jsonData))

	var personData Person
	error=json.Unmarshal(jsonData, &personData)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Person Data after unmarshaled:",personData)
}