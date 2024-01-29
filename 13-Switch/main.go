package main

import "fmt"

func main() {

	// Switch

	foo := 2

	switch {
	case foo == 1:
		println("bar")
	case foo == 2:
		println("buz")
	case foo == 3:
		println("fiz")
	default:
		println("foo")
	}

	var score float64
	println("Enter score: ")
	fmt.Scanf("%f", &score)
	switch {
	case score <= 59:
		println("F")
	case score <= 69:
		println("D")
	case score <= 79:
		println("C")
	case score <= 89:
		println("B")
	case score <= 100:
		println("A")
	default:
		println("Please enter a valid score.")
	}
}
