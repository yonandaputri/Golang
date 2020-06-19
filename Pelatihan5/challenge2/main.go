package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	number := 183928
	hasil := pasanganTerbesar(number)
	fmt.Println(hasil)
}

func pasanganTerbesar(angka int) int {
	numToString := strconv.Itoa(angka)
	sliceAngka := strings.Split(numToString, "")
	var max int
	for i := 0; i < len(sliceAngka); i++ {
		angka1, _ := strconv.Atoi(sliceAngka[i])
		puluhan := angka1 * 10
		for j := i + 1; j < len(sliceAngka); j++ {
			angka2, _ := strconv.Atoi(sliceAngka[i+1])
			tampungAngka := puluhan + angka2
			if tampungAngka > max {
				max = tampungAngka
			}
		}
	}
	return max
}
