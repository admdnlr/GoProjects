package main

import "fmt"

func main() {

	// Diziler

	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "Yellow"
	fmt.Println(colors[1])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Number of numbers:", len(numbers))

	myArray2 := [...]int{1, 2, 3, 4}
	fmt.Println(myArray2)

	var cars [3]string
	cars[0] = "Mercedes"
	cars[1] = "BMW"
	cars[2] = "Audi"
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}

	for j := 0; j < len(cars); j++ {
		fmt.Println(cars[j])
	}

	for i, value := range cars {
		fmt.Println("i = ", i, "value = ", value)
	}
}
