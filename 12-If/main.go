package main

import "fmt"

func main() {

	// IF

	a := 10
	b := 10

	if b > a {
		fmt.Println("b, a'dan büyüktür.")
	} else if b < a {
		fmt.Println("b, a'dan küçüktür.")
	} else {
		fmt.Println("b, a'ya eşittir.")
	}

	/*

		foo := 1
		if foo == 1 {
			fmt.Println("bar")
		}

	*/

	if foo := 2; foo == 1 {
		fmt.Println("bar")
	} else {
		fmt.Println("buz")
	}
}
