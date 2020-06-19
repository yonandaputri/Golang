//Viontina Dea IYP
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	polaPerkalian := "42#3 * 188 = 80#204"
	sliceString := strings.Split(perkalian, "=")
	kanan(sliceString[1])
	kiri(sliceString[0])
}

func tebakAngka(perkalian string) {
	hasilPerkalian := kiri(perkalian)
	pengali := kanan(perkalian)
	if hasilPerkalian == perkalian {

	}
}

func kiri(perkalian string) {
	kali := strings.Split(sliceString[0], "*")
	splitPalingKiri := strings.Split(kali[0], "")
	fmt.Println(kali)
	fmt.Println(hasilKali)
	for i := 0; i < len(splitPalingKiri); i++ {
		if splitPalingKiri[i] == "#" {
			splitPalingKiri[i] == "0"
		}
		convert, _ := strconv.Atoi(splitPalingKiri[i])
		ribuan := convert[0] * 1000
		ratusan := convert[1] * 100
		puluhan := convert[2] * 10
		bilanganKiri := ribuan + ratusan + puluhan + splitPalingKiri[3]
		for j := 0; j <= 9; i++ {

		}
	}
}

func kanan(hasil string) int {
	splitString := strings.Split(hasil, "")
	convert, _ := strconv.Atoi(hasil)
	for i := 0; i <= len(splitString); i++ {
		if splitString[i] == "#" {
			for j := 0; j <= 9; j++ {

			}
		}
	}
}
