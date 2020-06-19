// Viontina Dea Ivoni YP

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type randomNumber struct {
	number1, number2, number3 int
}

func main() {
	var totalCredit int
	creditBalance := 5000
	min := 1
	max := 9
	rand.Seed(time.Now().UnixNano())
	r := randomNumber{
		number1: rand.Intn(max-min+1) + min,
		number2: rand.Intn(max-min+1) + min,
		number3: rand.Intn(max-min+1) + min,
	}
	sum, check := r.tigaAngkaSama()
	sum2, check2 := r.duaAngkaSama()
	sum3 := r.tidakAdaAngkaSama()
	if check == true {
		fmt.Printf("You Win %d Dollars\n", sum)
		totalCredit = creditBalance + sum
		fmt.Printf("Your Total Credit Balance Is %d Dollar\n", totalCredit)
	} else if check2 == true {
		fmt.Printf("You Win %d Dollars\n", sum2)
		totalCredit = creditBalance + sum2
		fmt.Printf("Your Total Credit Balance Is %d Dollar\n", totalCredit)
	} else {
		fmt.Printf("You Lose %d Dollars\n", sum3)
		totalCredit = creditBalance - sum3
		fmt.Printf("Your Total Credit Balance Is %d Dollar\n", totalCredit)
	}
}

func (r *randomNumber) duaAngkaSama() (int, bool) {
	var jumlahAngka int
	var sliceNumber = []int{r.number1, r.number2, r.number3}
	var cek bool
	fmt.Println(sliceNumber)
	for i := 0; i < len(sliceNumber); i++ {
		for j := i + 1; j < len(sliceNumber); j++ {
			if sliceNumber[i] == sliceNumber[j] {
				jumlahAngka = (r.number1 + r.number2 + r.number3) * 100
				cek = true
			}
		}
	}
	return jumlahAngka, cek
}

func (r *randomNumber) tigaAngkaSama() (int, bool) {
	var jumlahAngka int
	var cek bool
	if r.number1 == r.number2 && r.number2 == r.number3 {
		jumlahAngka = (r.number1 + r.number2 + r.number3) * 200
		cek = true
	} else {
		cek = false
	}
	return jumlahAngka, cek
}

func (r *randomNumber) tidakAdaAngkaSama() int {
	jumlahAngka := (r.number1 + r.number2 + r.number3) * 50
	return jumlahAngka
}
