package main

import "fmt"

func main() {
	var angkaAwal, angkaAkhir int
	fmt.Print("Masukkan angka awal : ")
	fmt.Scan(&angkaAwal)
	fmt.Print("Masukkan angka akhir : ")
	fmt.Scan(&angkaAkhir)
	tampung, pesan := deretAngka(angkaAwal, angkaAkhir)
	fmt.Println(tampung, pesan)
	hasilAvg := rataRata(angkaAwal, angkaAkhir)
	fmt.Println(hasilAvg)
	jenis, hasilGenap, hasilGanjil := ganjilGenap(angkaAwal, angkaAkhir)
	fmt.Println(jenis)
	fmt.Print("Deret Genap : ")
	fmt.Println(hasilGenap)
	fmt.Print("Deret Ganjil : ")
	fmt.Print(hasilGanjil)
}

func deretAngka(awal, akhir int) ([]int, string) {
	var deret []int
	if awal > akhir {
		return deret, "\nInputan salah"
	} else {
		for i := awal; i <= akhir; i++ {
			deret = append(deret, i)
		}
		return deret, "\nInputan benar"
	}
}

func rataRata(awal, akhir int) float32 {
	var total int
	var deret []int
	for i := awal; i <= akhir; i++ {
		deret = append(deret, i)
		total += i
	}
	jumAngka := float32(len(deret))
	average := float32(total) / jumAngka
	return average
}

func ganjilGenap(awal, akhir int) (string, []int, []int) {
	var deret []int
	var ganjil, genap []int
	var jenisDeret string
	for i := awal; i <= akhir; i++ {
		deret = append(deret, i)
		if i%2 == 0 {
			//jenisDeret = "Genap"
			genap = append(genap, i)
		} else if i%2 != 0 {
			//jenisDeret = "Ganjil"
			ganjil = append(ganjil, i)
		} else {
			jenisDeret = "Angka tidak ada ganjil genap"
		}
	}
	return jenisDeret, genap, ganjil
}
