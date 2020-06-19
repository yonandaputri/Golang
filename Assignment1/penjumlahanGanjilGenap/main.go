//Viontina Dea IYP

package main

import "fmt"

func main() {
	// input
	var angka int
	var ganjil int
	var genap int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	fmt.Print("Ganjil = ")
	for i := 0; i < angka; i++ {
		if i%2 != 0 {
			fmt.Print(i)
			if i != 0 {
				fmt.Print(" + ")
			}
			ganjil += i
		}
	}
	fmt.Println(" =", ganjil)
	fmt.Print("Genap = ")
	for i := 0; i <= angka*2; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			if i != angka*2 {
				fmt.Print(" + ")
			}
			genap += i
		}
	}
	fmt.Println(" =", genap)
}
