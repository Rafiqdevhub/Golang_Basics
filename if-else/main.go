package main

func main() {
	// if-else
	age := 11
	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}

	// if-else if
	age = 6
	if age < 18 {
		println("You are a minor")
	} else if age < 60 {
		println("You are an adult")
	} else {
		println("You are a senior citizen")
	}
}