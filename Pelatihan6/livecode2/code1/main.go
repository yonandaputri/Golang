//Viontina Dea IYP
package main

import "fmt"

func main() {
	var coin = []int{1, 5, 7, 9, 11}
	var kombinasiAngka int
	fmt.Print("Masukkan jumlah kombinasi angka : ")
	fmt.Scan(&kombinasiAngka)
	if kombinasiAngka <= 33 {
		if kombinasiCoin2(coin, kombinasiAngka) == true {
			kombinasiCoin2(coin, kombinasiAngka)
		} else {
			kombinasiCoin3(coin, kombinasiAngka)
		}
	} else {
		fmt.Println("Kombinasi angka lebih dari 33")
	}
}

func kombinasiCoin2(koin []int, angka int) bool {
	//var kombinasi []int
	var bisa bool
	for i := 0; i < len(koin); i++ {
		for j := i + 1; j < len(koin); j++ {
			totalPenjumlahan := koin[i] + koin[j]
			if totalPenjumlahan == angka {
				bisa = true
				fmt.Printf("[%d %d]", koin[i], koin[j])
			} else {
				bisa = false
			}

		}
	}
	return bisa
}

func kombinasiCoin3(koin []int, angka int) {
	//var kombinasi []int
	for i := 0; i < len(koin); i++ {
		for j := i + 1; j < len(koin); j++ {
			for k := j + 1; k < len(koin); k++ {
				totalPenjumlahan := koin[i] + koin[j] + koin[k]
				if totalPenjumlahan == angka {
					fmt.Printf("[%d %d %d]", koin[i], koin[j], koin[k])
				}
			}
		}
	}
}

//1+5
