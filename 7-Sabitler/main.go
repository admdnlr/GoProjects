package main

import "fmt"

const thy = "Turkish Airlines"
const tai = "Turkish Aerospace Industries"

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	GOOGLE    Brand = "Google"
	MICROSOFT Brand = "Microsoft"
	DIJIBIL   Brand = "Dijibil"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {

	// const sabitlerin degeri program calisirken degistirilemez.
	fmt.Println(thy, tai) // Turkish Airlines

	PrintBrand(FACEBOOK) // Facebook
	PrintBrand(GOOGLE)   // Google
}
