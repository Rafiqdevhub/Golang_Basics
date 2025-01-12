package main

import "fmt"

func main() {
	// Create a struct
	// type Student struct {
	// 	Name  string
	// 	Grade int
	// }

	// Create a variable of type Student
	// var student Student

	// fmt.Println("Student:", student)

	// Assign values to the fields
	// student.Name = "Sana"
	// student.Grade = 90

	// student2 := Student{
	// 	Name:  "Ali",
	// 	Grade: 80,
	// }

	// var student3= new(Student)
	// student3.Name = "Baber"
	// student3.Grade = 85

	// Access the fields
	// fmt.Println("Student Name:", student.Name)
	// fmt.Println("Student Grade:", student.Grade)

	// fmt.Println("Student2 Name:", student2.Name)
	// fmt.Println("Student2 Grade:", student2.Grade)

	// fmt.Println("Student3 Name:", student3.Name)
	// fmt.Println("Student3 Grade:", student3.Grade)

	type Person struct {
		Name  string
		Age   int
		email string
	}

	type Contact struct {
		Phone   string
		Address string
	}

	type Employee struct {
		Person_detail  Person
		Person_contact Contact
	}

	var employee Employee

	employee.Person_detail = Person{
		Name:  "Ali",
		Age:   25,
		email: "ali@gmail.com",
		
	}

fmt.Println("Employee Name:", employee)
	fmt.Println("Employee Name:", employee.Person_detail.Name)
	fmt.Println("Employee Age:", employee.Person_detail.Age)
	fmt.Println("Employee Email:", employee.Person_detail.email)





}