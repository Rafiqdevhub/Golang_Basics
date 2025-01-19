package main

func modifyValue(num *int) {
	*num = *num + 10
}

func main() {

	// var num int
	num := 5

	// var ptr *int
	ptr := &num

	println("Value of num:", num)
	println("Address of num:", ptr)
	println("Data contain :", *ptr)

	value := 15
	modifyValue(&value)
	println("Value after modification:", value)
}