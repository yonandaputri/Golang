package main

import "fmt"

func main() {
	var angkaAwal, angkaAkhir int
	fmt.Print("Angka Awal : ")
	fmt.Scan(&angkaAwal)
	fmt.Print("Angka Akhir : ")
	fmt.Scan(&angkaAkhir)
	cariBilanganPrima(angkaAwal, angkaAkhir)
}

func cariBilanganPrima(bilanganAwal, bilanganAkhir int) {
	var hasil string
	for i := bilanganAwal; i <= bilanganAkhir; i++ {
		jumFaktor := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				jumFaktor++
			}
		}
		if jumFaktor == 2 {
			hasil = "bilangan prima"
		} else {
			hasil = "bukan bilangan prima"
		}
		fmt.Println(i, hasil)
	}
}

/*func main() {
	for i := 0; i <= 25; i++ {
		isPrima := true
		if i == 0 || i == 1 {
			fmt.Printf("%d bukan bilangan prima\n", i)
			continue
		} else {
			for j := i - 1; j > 1; j-- {
				if i%j == 0 {
					fmt.Printf("%d bukan bilangan prima\n", i)
					isPrima = false
					break
				}
			}
			if isPrima {
				fmt.Printf("%d bilangan prima\n", i)
			}
		}

	}
}*/
