package main

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// FOR RANGE

	// For range, bir dizi, slice, string veya map üzerinde dönmek için kullanılır.
	// For range, döngüdeki her bir eleman için bir kez çalışır.

	for i, v := range pow {
		println("2**", i, "=", v)
	}

	// [...] şu anlama geliyor: "Bu diziye ne kadar eleman eklersen ekle,
	// benim için önemi yok, benim için önemli olan bu diziye eleman
	// ekleyeceğini biliyorum, o yüzden senin yerine ben eleman
	// sayısını hesaplayacağım."

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		println("Array item", i, "is", a[i])
	}

	// MAP

	// Map'lerde de for range kullanabiliriz.
	// Map'lerde for range kullanırken, iki değer döndürür.
	// Birinci değer key, ikinci değer ise value'dur.

	baskentler := map[string]string{"Türkiye": "Ankara", "Fransa": "Paris", "Almanya": "Berlin", "İngiltere": "Londra"}

	for key, value := range baskentler {
		println("Map item: Başkent", value, "is", key)
	}
}
