package main

import "fmt"

func main() {
	umur := 35
	fmt.Println(golonganUmur(umur))
}
func golonganUmur(umur int) string {
	var hasil string
	if umur <= 5 {
		hasil = "Balita"
	} else if umur > 5 && umur <= 17 {
		hasil = "Anak-anak"
	} else if umur > 17 && umur <= 55 {
		hasil = "Dewasa"
	} else {
		hasil = "Orang Tua"
	}
	return hasil
}
