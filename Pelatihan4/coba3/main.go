package main

import (
	"fmt"
	"strings"
)

func main() {
	var dicari string
	var kalimat string = "aku"
	pisah := strings.Split(kalimat, "")
	fmt.Println(kalimat)
	fmt.Print("Input huruf yang dicari : ")
	fmt.Scan(&dicari)
	_, adaDi := cariSesuatu(pisah, dicari)
	fmt.Printf("Huruf ada di index : %d", adaDi)
}

func cariSesuatu(slice []string, huruf string) (hasil bool, index int) {
	//fmt.Println(slice)
	//var huruf string
	for i := 0; i < len(slice); i++ {
		if slice[i] == huruf {
			hasil = false
			index = -1
		} else {
			hasil = true
			index = i
		}
	}
	return hasil, index

}
