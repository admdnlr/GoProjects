package main

import "fmt"

var sayi int = 10

const aciklama = "Go programlama dili öğreniyorum"
const pi = 3.14

var a = "ABC XYZ"
var b, c = "ABC", "XYZ"
var d string

var (
	deigisken1 = "Adem"
	deigisken2 = "Dinler"
)

func main() {

	var message string
	message = "Hello, Go!"
	fmt.Println(message) // Hello, Go!

	var message2 = "Hello!"
	fmt.Println(message2) // Hello!

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c) // 1 2 3

	var d, e, f = 4.5, true, "Hello!"
	fmt.Println(d, e, f) // 4.5 true Hello!

	var g int
	var h bool
	var i string
	var j float64
	fmt.Println(g, h, i, j) // 0 false  0

	var k int
	var l, m string = "Hello", "World!"
	fmt.Println(k, l, m) // 0 Hello World!

	u := "Adem"
	fmt.Println(u) // Adem

	v, w, x := 2.1, false, "Adem"
	fmt.Println(v, w, x) // 2.1 false Adem

	y := "Metin"
	z := 'M'
	t := `Metin`
	fmt.Println(y, z, t) // Metin 77 Metin

	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	println(myFloat32) // +4.400000e+001
	println(myComplex) // (+3.000000e+000+4.000000e+000i)

	fmt.Println(sayi) // 10

	fmt.Println(aciklama) // Go programlama dili öğreniyorum

	fmt.Println(deigisken1) // Adem
}
