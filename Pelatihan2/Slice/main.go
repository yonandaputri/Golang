package main

import "fmt"

func main() {
	var kumpulanUmur = []int{}
	var kumpulanNama = []string{"budi", "tono", "tini"} // tidak perlu menambahkan jumlah yang akan dialokasi seperti array
	var kumpulanNama2 = kumpulanNama
	// bisa ditambahkan isinya lagi, lebih fleksible daripada array
	// slice hanya ada di golang, pengembangan dari array
	// array tidak bisa diubah jadi slice

	kumpulanUmur = append(kumpulanUmur, 7, 8, 9) // menambahkan element
	kumpulanUmur[0] = 10                         // menggantikan element yang sudah ada

	kumpulanNama2[1] = "mahmut" // kumpulanNama ikut berubah karena masih memakai memori yang sama pada slice
	// kalau pakai array kumpulanNama tidak akan ikut berubah karena alamat memorinya berbeda

	var kumpulan1 = [...]string{"a", "b", "c"} // array
	var kumpulan2 = kumpulan1[:]               // slice
	// [:] berarti data ditampilkan semua
	// depan untuk menghilangkan berapa kali dari depan
	// belakang untuk menghilangkan berdasarkan sampai index
	kumpulan2 = append(kumpulan2, "kumpulan1")
	fmt.Printf("%v\n", kumpulan2)

	var kumpulan3 = make([]string, 4) // len nya 4
	fmt.Println(kumpulan3)

	var kumpulan4 = kumpulan2[0:2] // mulai dari index 0 sebelum 2
	fmt.Println(len(kumpulan4))    // panjang variabel
	fmt.Println(cap(kumpulan4))    // kapasitas maksimum

	kumpulan4 = append(kumpulan4, "d", "e", "f", "g")
	fmt.Println(len(kumpulan4))
	fmt.Println(cap(kumpulan4))

	fmt.Println(kumpulanUmur)
	fmt.Println(kumpulanNama)
	fmt.Println(kumpulanNama2)

	// append slice menggunakan perulangan
	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)

}
