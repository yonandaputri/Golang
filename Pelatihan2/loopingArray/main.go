package main

import "fmt"

func main() {
	var kumpulan1 = [...]string{"budi", "tono", "tini", "sasa"} //array

	/*for i := len(kumpulan1) - 1; i >= 0; i-- {
		fmt.Println(kumpulan1[i])
	}*/
	// _ untuk ignore
	// range sama kaya
	// i untuk index
	// value untuk isinya
	for i, value := range kumpulan1 {
		fmt.Println(i, value)
	}
	// kalau ngga mau nampilin index, i dibuang atau diganti jadi _
	for _, value := range kumpulan1 {
		fmt.Println(value)
	}
}
