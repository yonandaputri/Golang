// Viontina Dea IYP

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(olehOleh(30000, []int{15000, 12000, 5000, 3000, 10000}))
	fmt.Println(olehOleh(10000, []int{2000, 2000, 3000, 1000, 2000, 10000}))
	fmt.Println(olehOleh(4000, []int{7500, 1500, 2000, 3000}))
	fmt.Println(olehOleh(50000, []int{25000, 25000, 10000, 15000}))
	fmt.Println(olehOleh(0, []int{10000, 3000}))
}

func olehOleh(maxBudget int, gifts []int) int {
	sort.Ints(gifts)
	jumlahBarang := 0
	for i, _ := range gifts {
		if gifts[i] <= maxBudget {
			maxBudget -= gifts[i]
			jumlahBarang++
		}
	}
	return jumlahBarang
}
