// https://dasarpemrogramangolang.novalagung.com/27-interface-kosong.html

package main

import (
	"fmt"
)

type apaAja interface {
}

type buku struct {
	judulBuku string
}

func (b buku) cetakJudul() string {
	return fmt.Sprintf("Judulnya :%s", b.judulBuku)
}

func main() {
	// var angkaPeaak apaaja = 100
	// var tulisanPeaakk apaaja = "String"
	// var isPeakKawe = true

	// var daftarVarPeak = []apaaja{angkaPeaak, tulisanPeaakk, isPeakKawe}
	// for _, a := range daftarVarPeak {
	// 	printPea(a)
	// }
	// var b = buku{"Anny Arrow"}
	// printPea(b)
	// myPrintLen(1, "satu", true, []int{})
}

func printPea(a apaAja) {
	// Menggunakan Switch Case Type
	switch val := a.(type) {
	case buku:
		fmt.Println(val.cetakJudul())
	}
	// Melakukan Casting
	// fmt.Println(a.(buku).cetakJudul())
}

// func myPrintLen(a ...interface{}) {
// 	fmt.Print(a...)
// }
