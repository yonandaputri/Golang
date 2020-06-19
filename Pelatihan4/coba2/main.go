package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "i will become a full stack developer"
	hasilHitung := hitungKata(kata)
	fmt.Println("Jumlah kata : ", hasilHitung)
	hasilHitungVokal := hitungVokal(kata)
	fmt.Println("Jumlah huruf Vokal : ", hasilHitungVokal)
}

func hitungKata(kalimat string) int {
	sliceKata := strings.Split(kalimat, " ")
	jumKata := len(sliceKata)
	return jumKata
}

func hitungVokal(kalimat string) int {
	var jumVokal int
	huruf := strings.Split(kalimat, "")
	for _, alphabet := range huruf {
		if alphabet == "a" || alphabet == "i" || alphabet == "u" || alphabet == "e" || alphabet == "o" {
			jumVokal++
		}
	}
	return jumVokal
}
