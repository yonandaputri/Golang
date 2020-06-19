package main

import "fmt"

func main() {
	// array = kumpulan data
	// var kumpulanNama = [3]string{"antony", "garin", "ben"} // cara input pertama
	var kumpulanUmur [3]int // cara input kedua
	kumpulanUmur[0] = 10
	kumpulanUmur[1] = 15
	kumpulanUmur[2] = 20
	// assign nilai tidak boleh melebihi panjang arraynya
	// index dimulai dari 0
	kumpulanNama := make([]string, 10)          // 10 itu panjang arraynya
	kumpulanNama = append(kumpulanNama, "tono") // karena append jadi ditambahkan diujung

	x := [...]int{10, 20, 30}         // automatic range len menggunakan ...
	x[2] = 100                        // bisa assign by index
	var y [5]int = [5]int{10, 20, 30} // index yang tidak diisi jadi 0 karena harus ada 5 jumlahnya, valuenya harus ada tidak boleh ...
	z := [5]int{1: 10, 3: 30}         //menginisialisasi array menggunakan indexnya i=1 jadi 10, i=3 jadi 30

	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(len(kumpulanUmur))
	fmt.Println(kumpulanNama)
	fmt.Println(kumpulanUmur)
}
