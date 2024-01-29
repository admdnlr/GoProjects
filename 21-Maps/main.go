package main

import "fmt"

func main() {

	// Maps

	// Map'ler key-value şeklinde çalışır.
	// Key'ler unique'dir.
	// Key'lerin tipi aynı olmalıdır.
	// Key'lerin tipi int, float, string, bool olabilir.
	// Value'ların tipi aynı olmalıdır.
	// Value'ların tipi int, float, string, bool, array, slice, map olabilir.
	// Map'ler referans tip olduğu için, fonksiyonlara parametre olarak gönderilirken dikkat edilmelidir.

	// Map oluşturma
	// 1. Yöntem
	var myMap1 map[string]int
	myMap1 = make(map[string]int)
	myMap1["key1"] = 123
	myMap1["key2"] = 234
	myMap1["key3"] = 345
	myMap1["key4"] = 456

	// 2. Yöntem
	myMap2 := make(map[string]int)
	myMap2["key1"] = 123
	myMap2["key2"] = 234
	myMap2["key3"] = 345

	// 3. Yöntem
	myMap3 := map[string]int{
		"key1": 123,
		"key2": 234,
		"key3": 345,

		// Key'ler unique olduğu için, aynı key ile bir daha ekleme yaparsak hata alırız.
		// "key3": 456,
	}
	myMap3["key4"] = 456

	myMap := make(map[int]string)

	fmt.Println(myMap)
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)

	states := make(map[string]string)
	states["IST"] = "Istanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	antalya := states["ANT"]
	fmt.Println("seçili şehir : ", antalya)

	delete(states, "ANT")
	fmt.Println(states)

	states["IZM"] = "Izmir"
	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
