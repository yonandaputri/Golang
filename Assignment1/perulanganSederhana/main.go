//Viontina Dea IYP

package main

import "fmt"

func main() {
	// input
	var kalimat string = "I will become a go developer"
	var angka int
	fmt.Print("Masukkan jumlah perulangan : ")
	fmt.Scan(&angka)

	for i := angka; i > 0; i-- {
		fmt.Println(i, "-", kalimat)
	}
}
