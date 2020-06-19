package main

import "fmt"

// type data buatan, string jadi stringku
type stringku string

/*var penjumlah = func(a int, b int) int {
	return a + b
}

var pengurangan = func(a int, b int) int {
	return a - b
}*/

// func bisa dijadikan parameter
// bisa di dalam main atau di luar main
func main() {
	//var s stringku = "String baru"
	//cetak(penjumlah, 1, 2)
	//cetak(pengurangan, 4, 2)
	angka1 := 10
	angka2 := 5
	pengaliKu := 100
	// anonymous func parameternya dibawah
	func(pengali int) {
		var jumlah int
		jumlah = (angka1 + angka2) * pengali
		fmt.Println(jumlah)
	}(pengaliKu)

	cetakBanyak("a", "b", "c", "d")

	// defer dieksekusi terakhiran setelah semua function selesai dieksekusi
	// untuk tutup database atau file. jadi closing
	kedua()
	pertama()
	defer func() {
		fmt.Println("Penutup")
	}()

	// defer itu LIFO
	// kalau defer ada 2 atau lebih, defer yang terakhir yang dieksekusi duluan

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

// menerima 3 parameter
func cetak(ops func(int, int) int, a, b int) {
	hasil := ops(a, b)
	fmt.Println(hasil)
}

// variadic function
// walau typedatanya string tetap jadi slice hasilnya
func cetakBanyak(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// riset tentang high order function
// func bisa mereturn sebuah func

// deferred function
// untuk menutup file yang sudah dibuka dan diedit
func pertama() {
	fmt.Println("pertama")
}

func kedua() {
	fmt.Println("kedua")
}
