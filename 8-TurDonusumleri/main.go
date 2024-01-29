package main

import (
	"fmt"
	"strconv"
)

func main() {

	var myString string = "1"
	var x = 10
	var f float32 = 2.5

	fmt.Println(myString, x, f) // 1 10 2.5

	// Stringten int'e donusum
	number, _ := strconv.Atoi(myString)
	fmt.Println(number) // 1

	result := number + 2
	fmt.Println(result) // 3

	fmt.Println(myString + strconv.Itoa(result)) // 13

	// casting (tip donusumu) islemi yapilirken dikkatli olunmalidir.
	// veri kaybi olabilir.
	var i int = 55
	var j float64 = float64(i)
	var k uint = uint(j)
	fmt.Println(i, j, k) // 55 55 55

}
