package main

func main() {
	// switch
	age := 19
	switch {
	case age >= 18:
		println("You can vote")
	default:
		println("You are not eligible to vote")
	}

	// switch with expression
	age = 00
	switch {
	case age < 18 && age >= 0:
		println("You are a minor")
	case age < 60:
		println("You are an adult")
	default:
		println("You are a senior citizen")
	}
}