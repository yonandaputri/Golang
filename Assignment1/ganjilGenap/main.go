//Viontina Dea IYP

package main

import "fmt"

func main() {
	// input
	var angka int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}
}
