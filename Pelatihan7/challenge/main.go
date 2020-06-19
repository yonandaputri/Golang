// Viontina Dea Ivoni YP

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var dataMahasiswa []mahasiswa

type mahasiswa struct {
	nama, jurusan string
	umur          int
}

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	mainMenu()
}

func mainMenu() {
	fmt.Println("--------------------------------------")
	fmt.Println("Main Menu")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Keluar")
	fmt.Println("Pilihan menu dari 1-4")
}

func valPanjangNama(nama string, panjangNamaMin, panjangNamaMax int) bool {
	splitNama := strings.Split(nama, "")
	if len(splitNama) < panjangNamaMin && len(splitNama) > panjangNamaMax {
		return false
	} else {
		return true
	}
}

func valUmur(umur, umurMin int) bool {
	if umur < umurMin {
		return false
	} else {
		return true
	}
}

func valPanjangJurusan(jurusan string, panjangJurusanMax int) bool {
	splitJurusan := strings.Split(jurusan, "")
	if len(splitJurusan) <= 10 {
		return true
	} else {
		return false
	}
}

func addMahasiswa() {
	var konfirmasi string
	var namaMahasiswa string
	var umurMahasiswa int
	var jurusanMahasiswa string

	fmt.Println("--------------------------------------")
	fmt.Println("Add Mahasiswa")
	fmt.Println("--------------------------------------")
	fmt.Print("Nama (3-20 Karakter) : ")
	fmt.Scanln(&namaMahasiswa)
	fmt.Print("Umur (min 17 tahun) : ")
	fmt.Scanln(&umurMahasiswa)
	fmt.Print("Jurusan (maks 10 Karakter) : ")
	fmt.Scanln(&jurusanMahasiswa)

	isValidNama := valPanjangNama(namaMahasiswa, 3, 20)
	isValidUmur := valUmur(umurMahasiswa, 17)
	isValidJurusan := valPanjangJurusan(jurusanMahasiswa, 10)

	if isValidNama == true && isValidUmur == true && isValidJurusan == true {
		fmt.Printf("Apakah anda yakin akan menyimpan data tersebut (y/n) ? ")
		fmt.Scanln(&konfirmasi)
		if strings.ToLower(konfirmasi) == "y" && len(dataMahasiswa) <= 5 {
			mhsBaru := mahasiswa{nama: namaMahasiswa, umur: umurMahasiswa, jurusan: jurusanMahasiswa}
			dataMahasiswa = append(dataMahasiswa, mhsBaru)
			clearScreen()
		} else if strings.ToLower(konfirmasi) == "n" {
			fmt.Println("Data Tidak Tersimpan!")
			time.Sleep(time.Millisecond * 1500)
			clearScreen()
		} else if len(dataMahasiswa) > 5 {
			fmt.Println("Data Mahasiswa sudah penuh!")
			time.Sleep(time.Millisecond * 1500)
			clearScreen()
		}
	} else {
		fmt.Println("Data Yang Dimasukkan Tidak Sesuai! Data Tidak Tersimpan!")
		time.Sleep(time.Millisecond * 1500)
		clearScreen()
	}
}

func delMahasiswa() {
	if len(dataMahasiswa) < 1 {
		fmt.Println("--------------------------------------")
		fmt.Println("Data kosong")
		fmt.Println("--------------------------------------")
		time.Sleep(time.Millisecond * 1500)
		clearScreen()
	} else {
		dataMahasiswa = dataMahasiswa[:len(dataMahasiswa)-1]
		fmt.Println("--------------------------------------")
		fmt.Println("Data terakhir telah dihapus")
		fmt.Println("--------------------------------------")
		time.Sleep(time.Millisecond * 1500)
		clearScreen()
	}
}

func lihatMahasiswa() {
	var pilihMenu int
	fmt.Println("--------------------------------------")
	fmt.Println("View Mahasiswa")
	fmt.Println("--------------------------------------")
	fmt.Println("1. View By Index")
	fmt.Println("2. View All")
	fmt.Print("Masukan menu yang dipilih : ")
	fmt.Scan(&pilihMenu)
	if pilihMenu == 1 {
		lihatByIndex()
	} else if pilihMenu == 2 {
		lihatSemua()
	}
}

func lihatByIndex() {
	var pilihIndex int
	fmt.Println("--------------------------------------")
	fmt.Println("View By Index")
	fmt.Println("--------------------------------------")
	fmt.Print("index yang ingin ditampilkan : ")
	fmt.Scan(&pilihIndex)
	if len(dataMahasiswa) < 1 {
		fmt.Println("Data Kosong!!!")
	} else {
		for i, data := range dataMahasiswa {
			if pilihIndex == i {
				fmt.Println(i)
				fmt.Println("Nama : ", data.nama)
				fmt.Println("Umur : ", data.umur)
				fmt.Println("Jurusan : ", data.jurusan)
			}
		}
	}
	fmt.Println("--------------------------------------")
	time.Sleep(time.Millisecond * 1500)
	mainMenu()
}

func lihatSemua() {
	fmt.Println("--------------------------------------")
	fmt.Println("View All")
	fmt.Println("--------------------------------------")
	if len(dataMahasiswa) < 1 {
		fmt.Println("Data Kosong!!!")
	} else {
		for i, data := range dataMahasiswa {
			fmt.Println(i)
			fmt.Println("Nama : ", data.nama)
			fmt.Println("Umur : ", data.umur)
			fmt.Println("Jurusan : ", data.jurusan)
		}
	}
	fmt.Println("--------------------------------------")
	time.Sleep(time.Millisecond * 1500)
	mainMenu()
}

func main() {
	mainMenu()
	for {
		var selectedMenu string

		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			addMahasiswa()
			break
		case "2":
			delMahasiswa()
			break
		case "3":
			lihatMahasiswa()
			break
		case "4":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
