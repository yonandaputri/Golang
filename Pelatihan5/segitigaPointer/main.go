package main

import "fmt"

func main() {
	var luas float64
	alas := 10.0
	tinggi := 12.0
	luasSegitiga(&luas, &alas, &tinggi)
	fmt.Println(luas)
}

func luasSegitiga(luas, alas, tinggi *float64) {
	*luas = *alas * *tinggi * 0.5
}
