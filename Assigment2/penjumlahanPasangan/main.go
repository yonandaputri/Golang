// Viontina Dea IYP

package main

import "fmt"

func main() {
	var number = []int{3, 5, 2, -4, 8, 11}
	fmt.Println(number)
	var sum int
	fmt.Print("Masukkan sum : ")
	fmt.Scan(&sum)

	fmt.Print("[")
	penjumlahanPasangan(number, sum)
	fmt.Print("]")

}

func penjumlahanPasangan(angka []int, jumlah int) {
	for i, _ := range angka {
		for j := i + 1; j < len(angka); j++ {
			totalPenjumlahan := angka[i] + angka[j]
			if totalPenjumlahan == jumlah {
				fmt.Printf("[%d %d]", angka[i], angka[j])
			}
		}
	}
}
