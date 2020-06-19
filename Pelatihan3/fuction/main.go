package main

import "fmt"

func main() {
	//var name = "budi"
	//var name2 = "tono"
	age := 20

	var umur = greet(age)
	fmt.Println(umur)

	namaFunction(age, 15, "budi")
	//fmt.Println(greet(name))
	/*greet(name)
	greet(name2)*/

	// cetak salah satu harus di ignore satunya
	hasilProsesLuas, _ := hitungLuasSegitiga(4, 8)
	fmt.Printf("Luas Segitiga Adalah %d\n", hasilProsesLuas)

	input := ganjilGenap(3)
	if input {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}

// string yang didalam kurung untuk type data return, bisa lebih dari 1
// misal (string, int)
func greet(umurDiCetak int) int {
	//fmt.Printf("Selamat Datang %s\n", nama)
	//return nama + " Tambahan"
	fmt.Println("Umur jadi : ")
	return umurDiCetak + 20
}

func namaFunction(age, size int, name string) {
	fmt.Println("Function Dipanggil", age, size, name)
}

func hitungLuasSegitiga(alas, tinggi int) (luasnya, volumenya int) {
	luas := (alas * tinggi) / 2
	volume := alas * tinggi

	return luas, volume
}

func ganjilGenap(angka int) bool {
	var hasil bool
	if angka%2 == 0 {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}
