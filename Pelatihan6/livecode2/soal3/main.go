/*1. Jika angka tidak habis dibagi 2 atau 3, maka function akan
me-return "SALAH".
2. Jika angka habis dibagi 2, maka return "TES"
3. Jika angka habis dibagi 3, maka return "LA"
4. Jika angka habis dibagi 2 dan 3, maka return "TESLA"*/

// Viontina Dea Ivoni YP
package main

import "fmt"

func main() {
	num := 4
	var cetak string
	habisDibagi(&cetak, &num)
	fmt.Println(cetak)
}

func habisDibagi(hasil *string, angka *int) {
	if *angka%2 == 0 && *angka%3 != 0 {
		*hasil = "TES"
	} else if *angka%3 == 0 && *angka%2 != 0 {
		*hasil = "LA"
	} else if *angka%2 == 0 && *angka%3 == 0 {
		*hasil = "TESLA"
	} else if *angka%2 != 0 && *angka%3 != 0 {
		*hasil = "SALAH"
	}
}
