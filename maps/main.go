package main

import "fmt"

func main() {
	studentsGrade := make(map[string]int)

	studentsGrade["John"] = 90
	studentsGrade["Alice"] = 80
	studentsGrade["Bob"] = 85
	studentsGrade["Charlie"] = 95

	fmt.Println("Students Grade:", studentsGrade["John"])
	studentsGrade["John"] = 95
	fmt.Println("Students new Grade:", studentsGrade["John"])

	delete(studentsGrade, "John")
	// fmt.Println("Students Grade after delete:", studentsGrade["John"])

	// check if key exists
	grade, ok := studentsGrade["Bob"]
	if ok {
		fmt.Println("John's Grade:", grade)
	} else {
		fmt.Println("John's Grade not found")
	}

	// iterate over map
	for key, value := range studentsGrade {
		fmt.Println("Key:", key, "Value:", value)
	}
}