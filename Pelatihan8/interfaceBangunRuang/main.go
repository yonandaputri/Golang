// sebenernya interface tuh method di dalam struct

package main

import "fmt"

type bangunDatar interface {
	GetLuas() float64
	GetKeliling() float64
}

type lingkaran struct {
	jariJari float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) GetLuas() float64 {
	return 3.14 * l.jariJari * l.jariJari
}

func (p persegi) GetLuas() float64 {
	return p.sisi * p.sisi
}

func (l lingkaran) GetKeliling() float64 {
	return 2 * 3.14 * l.jariJari
}

func (p persegi) GetKeliling() float64 {
	return 4 * p.sisi
}

func main() {
	hitungLingkaran := lingkaran{3}
	hitungPersegi := persegi{4}
	var inputan int
	fmt.Print("Input pilihan cetak : ")
	fmt.Scan(&inputan)
	fmt.Println("Hasil Hitung Lingkaran")
	printHasilHitung(hitungLingkaran, inputan)
	fmt.Println("Hasil Hitung Persegi")
	printHasilHitung(hitungPersegi, inputan)
}

func printHasilHitung(b bangunDatar, input int) {
	switch input {
	case 1:
		fmt.Println("Luas : ", b.GetLuas())
	case 2:
		fmt.Println("Keliling : ", b.GetKeliling())
	default:
		fmt.Println("Inputan Tidak Sesuai")
	}
}
