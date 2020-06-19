package main

import (
	"fmt"
)

func main() {
	var panjang int
	var lebar int
	var tinggi int

	// input
	fmt.Print("Masukkan panjang : ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar : ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&tinggi)

	// rumus
	var luasPermukaan int = 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	var luasVolume int = panjang * lebar * tinggi

	// print hasil
	//P = 6, L = 3, T = 4
	fmt.Println("P =", panjang, ", L =", lebar, ", T =", tinggi)
	fmt.Println("Luas Permukaan :", luasPermukaan)
	fmt.Println("Luas Volume :", luasVolume)
}
