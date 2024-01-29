package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Break

	fmt.Println("Break kullanımı:")
	i := 0
	for {
		fmt.Println("i'nin değeri : " + strconv.Itoa(i))
		i++
		if i > 10 {
			break
		}
	}
	fmt.Println("")

	fmt.Println("Continue kullanımı:")
	// Continue

	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("çikti : ", i)
	}
	fmt.Println("işleme devam ediyor...")
}
