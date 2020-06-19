package main

import (
	"fmt"
)

func main() {
	var bilangan = []int{1, 10, 45, 67, 33, 99, 2, 8}
	min, max, avg := minMaxAverage(bilangan)
	fmt.Println(bilangan)
	fmt.Printf("Hasil minimal : %d\n", min)
	fmt.Printf("Hasil maximal : %d\n", max)
	fmt.Printf("Hasil average : %f\n", avg)
}

func minMaxAverage(bil []int) (min, max int, avg float32) {
	var minimal int
	var maksimal int
	var jumlah float32
	for i := 0; i < len(bil); i++ {
		for j := i + 1; j < len(bil); j++ {
			if bil[i] < bil[j] {
				minimal = bil[i]
			} else if bil[i] > bil[j] {
				maksimal = bil[i]
			}
		}
		jumlah += float32(bil[i])
	}
	average := jumlah / float32(len(bil))
	return minimal, maksimal, average
}
