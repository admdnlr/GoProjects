package main

import "fmt"

func main() {

	a := 10
	b := 20

	total := a + b
	fmt.Println(total) // 30

	total = total - 5
	fmt.Println(total) // 25

	total *= 20
	fmt.Println(total) // 500

	total = 10 / 2
	fmt.Println(total) // 5

	total = a % b
	fmt.Println(total) // 10

	total++
	fmt.Println(total) // 11

	total--
	fmt.Println(total) // 10

}
