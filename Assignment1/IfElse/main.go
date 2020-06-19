// Viontina Dea Ivoni YP

package main

import "fmt"

func main() {
	var nama string = "Saitama"
	var peran string = "superhero"

	// kode disini
	if nama == "" && peran == "" {
		fmt.Println("Nama dan Peran Harus Diisi")
	} else if nama != "" && peran == "" {
		fmt.Println("Peran Harus Diisi")
	} else if nama != "" && peran == "superhero" {
		fmt.Println("Selamat Datang Superhero", nama, ", Kalahkan Semua Monster Di Muka Bumi")
	} else if nama != "" && peran == "monster" {
		fmt.Println("Selamat Datang Monster", nama, ", Hancurkan Semua Superhero Yang Ada")
	} else if nama != "" && peran != "superhero" || peran != "monster" {
		fmt.Println("Selamat Datang", nama, ", Pilih Peranmu Untuk Melanjutkan Game Ini")
	}
}
