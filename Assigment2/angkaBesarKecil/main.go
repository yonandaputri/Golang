// Viontina Dea IYP

package main

import "fmt"

func main() {
	var arr1 = []int{3, 9, 10, 13, 1, 2}
	min, max := minMax(arr1)
	fmt.Println(arr1)
	fmt.Println("Terkecil : ", min)
	fmt.Println("Terbesar : ", max)

}

func minMax(bil []int) (min, max int) {
	var minimal int
	var maksimal int
	for i, _ := range bil {
		for j := i + 1; j < len(bil); j++ {
			if bil[i] < bil[j] {
				minimal = bil[i]
			} else if bil[i] > bil[j] {
				maksimal = bil[i]
			}
		}
	}
	return minimal, maksimal
}
