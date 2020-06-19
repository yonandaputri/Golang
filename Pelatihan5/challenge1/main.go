package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var productList []map[string]string
var categoriList []map[string]string

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
	fmt.Println("**************************************")
	fmt.Println("Aplikasi")
	fmt.Println("**************************************")
	fmt.Println("1. Buat Produk Baru")
	fmt.Println("2. Daftar Produk")
	fmt.Println("3. Buat Kategori Baru")
	fmt.Println("4. Daftar Kategori")
	fmt.Println("5. Keluar")
	fmt.Println("Pilihan menu dari 1-5")
}

func listProductForm() {
	fmt.Println("Halaman Daftar Produk")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Produk", "Nama Produk")
	for idx, product := range productList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, product["productCode"], product["productName"])
	}
}

func newProductForm() {
	var productCode string
	var productName string
	var saveProductConfirmation string

	fmt.Println("Halaman Buat Produk Baru")
	fmt.Println("**************************************")
	fmt.Print("Product Code : ")
	fmt.Scanln(&productCode)
	fmt.Print("Product Name : ")
	fmt.Scanln(&productName)
	fmt.Printf("Produk %s dengan kode %s akan disimpan (y/n) ? ", productName, productCode)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "y" {
		newProduct := make(map[string]string)
		newProduct["productCode"] = productCode
		newProduct["productName"] = productName
		productList = append(productList, newProduct)
		clearScreen()
	} else if saveProductConfirmation == "n" {
		clearScreen()
	}
}

func listCategoriForm() {
	fmt.Println("Halaman Daftar Categori")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Categori", "Nama Categori")
	for index, categori := range categoriList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", index, categori["categoriCode"], categori["categoriName"])
	}
}

func newCategoriForm() {
	var categoriCode string
	var categoriName string
	var saveCategoriConfirmation string

	fmt.Println("Halaman Buat Categori Baru")
	fmt.Println("**************************************")
	fmt.Print("Categori Code : ")
	fmt.Scanln(&categoriCode)
	fmt.Print("Categori Name : ")
	fmt.Scanln(&categoriName)
	fmt.Printf("Categori %s dengan kode %s akan disimpan (y/n) ? ", categoriName, categoriCode)
	fmt.Scanln(&saveCategoriConfirmation)

	if saveCategoriConfirmation == "y" {
		newCategori := make(map[string]string)
		newCategori["categoriCode"] = categoriCode
		newCategori["categoriName"] = categoriName
		categoriList = append(categoriList, newCategori)
		clearScreen()
	} else if saveCategoriConfirmation == "n" {
		clearScreen()
	}
}

func main() {
	mainMenu()
	for {
		var selectedMenu string

		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			newProductForm()
			break
		case "2":
			listProductForm()
			break
		case "3":
			newCategoriForm()
			break
		case "4":
			listCategoriForm()
			break
		case "5":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
