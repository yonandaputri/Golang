package main

import (
	"Pelatihan8/interfaceCalculation/calculator"
	"fmt"
)

func main() {
	var input1, input2 int
	fmt.Print("Input angka 1 : ")
	fmt.Scan(&input1)
	fmt.Print("Input angka 2 : ")
	fmt.Scan(&input2)
	fmt.Println("Hasil Pengurangan")
	substraction := calculator.Pengurangan{input1, input2}
	printHasilHitung(substraction)
	fmt.Println("Hasil Penjumlahan")
	addition := calculator.Penjumlahan{input1, input2}
	printHasilHitung(addition)
}

func printHasilHitung(c calculator.Calculator) {
	hasil, pesanError := c.GetCalculation()
	if pesanError == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(pesanError.Error())
	}

}
