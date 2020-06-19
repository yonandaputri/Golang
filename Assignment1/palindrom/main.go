//Viontina Dea IYP

package main

import "fmt"

func main() {
	// input
	var kata string
	var hasil bool
	fmt.Print("Input kata : ")
	fmt.Scan(&kata)

	for i := 0; i < len(kata)/2; i++ {
		if kata[i] == kata[len(kata)-i-1] {
			hasil = true
		} else {
			hasil = false
		}
	}
	fmt.Println(hasil)
}
