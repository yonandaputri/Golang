// Viontina Dea IYP

package main

import "fmt"

func main() {
	var number = []int{1, 2, 3, 4, 5}
	fmt.Println(number)
	var sum int
	fmt.Print("Target Nilai Penjumlahan : ")
	fmt.Scan(&sum)

	hitungJumlahPasangan(number, sum)
}

func hitungJumlahPasangan(angka []int, jumlah int) {
	var jumlahPasangan int
	for i, _ := range angka {
		for j := i + 1; j < len(angka); j++ {
			totalPenjumlahan := angka[i] + angka[j]
			if totalPenjumlahan == jumlah {
				jumlahPasangan++
				fmt.Printf("[%d %d]\n", angka[i], angka[j])
			}
		}
	}
	fmt.Println("Jumlah Pasangan Penjumlahan : ", jumlahPasangan)
}
