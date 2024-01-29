package main

import (
	"fmt"
	"time"
)

func main() {

	// Konsol : Saat ve Tarih

	// Date() fonksiyonu ile tarih bilgisi alinir.
	t := time.Date(2020, time.July, 10, 20, 0, 0, 0, time.UTC)

	// t adıyla oluşturulan değişkenin tarih bilgisi ekrana yazdırılır.
	fmt.Printf("Go launched at %s\n", t)

	fmt.Println("---------------------------")

	// time.Now() fonksiyonu ile anlık tarih bilgisi alınır.
	now := time.Now()

	// now adıyla oluşturulan değişkenin tarih bilgisi ekrana yazdırılır.
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("---------------------------")

	// ilk başta oluşturulan t adlı değişkenin tarih bilgisi ekrana yazdırılır.
	fmt.Println("The month is:", t.Month())
	fmt.Println("The day is:", t.Day())
	fmt.Println("The weekday is:", t.Weekday())

	fmt.Println("---------------------------")

	// tarihe 1 gün eklenir.
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	fmt.Println("---------------------------")

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	fmt.Println("---------------------------")

	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
