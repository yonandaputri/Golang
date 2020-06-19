package main

import "fmt"

func main() {
	var karyawan = map[int]string{
		1111: "Anang",
		2222: "Mamat",
		3333: "Ucil",
	}
	for key, value := range karyawan {
		fmt.Println("NIK : ", key, "Nama karyawan : ", value)
	}
	var nik int
	fmt.Print("Input NIK Karyawan yang dicari : ")
	fmt.Scan(&nik)

	var nama, isExist = karyawan[nik]
	if isExist {
		fmt.Println("Nama Karyawan : ", nama)
	} else {
		fmt.Println("Nama karyawan yang dicari tidak ada")
	}

}
