package main

import "fmt"

func main() {

	// _ bos tanimlayici olarak kullanilir.
	// Bir degiskenin degerini atamak istemedigimizde kullanilir.
	// Bu degiskenin degeri kullanilmayacagi icin hata vermez.

	var _, x, _ int = 1, 2, 3
	fmt.Println(x) // 2
}
