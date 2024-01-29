package main

import (
	"fmt"
)

func main() {
	// Slice

	// Slice, dizilerin bir parçasını almak için kullanılır.
	// Dizilerin boyutu sabit olduğu için slice'lar kullanılır.
	// Slice'lar dizilerin referansını tutar.
	// Slice'lar dizilerin bir parçasını alırken, dizilerin bir parçasını değiştirirken kullanılır.

	primes := []int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // 1. index'ten 4. index'e kadar olan elemanları alır.
	fmt.Println(s)

	// Slice'lar referans tip olduğu için, bir slice'ın elemanlarını değiştirirsek,
	// slice'ın referans gösterdiği dizinin elemanları da değişir.

	myArray1 := [...]int{45, 23, 43}
	mySlice1 := myArray1[:]
	fmt.Println(mySlice1)
	mySlice1[0] = 99
	fmt.Println(mySlice1)
	fmt.Println(myArray1)

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 234
	numbers[2] = 345
	numbers[3] = 456
	numbers[4] = 567
	fmt.Println(numbers)

	numbers = append(numbers, 678)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 789)
	fmt.Println(cap(numbers))
	numbers = append(numbers, 890)
	numbers = append(numbers, 901)
	fmt.Println(cap(numbers))
	numbers = append(numbers, 12)
	fmt.Println(cap(numbers))
	numbers = append(numbers, 23)
	fmt.Println(cap(numbers)) // kapasite 16'dan 32'ye çıktı.
}
