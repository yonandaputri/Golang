package main

import "fmt"

// map bentuknya key dan value
func main() {
	// buat nyimpen, key sama kaya index kalau di array
	// key bersifat fleksible
	var names = map[string]bool{}
	// yang didalem kurung kotak buat type data key
	// yang diluar buat type data value

	// untuk inisialisasi
	names["name"] = true
	names["age"] = false

	var warna = make(map[string]string)
	var warnaKey string = "merah"
	var warnaValue string = "FF0000"
	warna[warnaKey] = warnaValue
	warna["hijau"] = "FFFFFF"

	// isExist bisa diganti apa aja
	// _ itu value
	var _, isExist = warna["merah"]

	if isExist {
		fmt.Println("ADA")
	} else {
		fmt.Println("ENGGA ADA")
	}

	// looping
	// kalau di array kunci untuk index
	// isi untuk value
	// berarti index
	for kunci, isi := range warna {
		fmt.Println(kunci, isi)
	}
	fmt.Println(len(warna))

	fmt.Println(names)
	fmt.Println(warna["hijau"])

	// delete map dengan key
	delete(warna, "merah")
	fmt.Println(warna)

	// assign banyak key dan value pada map
	var tanggal = map[string]int{
		"januari":  01,
		"februari": 02,
		"maret":    03,
	}
	for key, val := range tanggal {
		fmt.Println(key, val)
	}
	fmt.Println(len(tanggal))
}
