// Viontina Dea Ivoni YP

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 = []string{"-1", "-1", "-1", "-1", "-1", "-2"}
	var hasil1 int
	var input2 = []string{"+5", "-3", "+5", "+5", "-1"}
	var hasil2 int
	var input3 = []string{"+4", "+3", "+7", "+1"}
	var hasil3 int

	//fmt.Print("Input serangan berupa angka : ")
	//fmt.Scan(&angka)

	for i := 0; i < len(input1); i++ {
		angka1, _ := strconv.Atoi(input1[i])
		hasil1 += angka1
	}
	for i := 0; i < len(input2); i++ {
		angka2, _ := strconv.Atoi(input2[i])
		hasil2 += angka2
	}
	for i := 0; i < len(input3); i++ {
		angka3, _ := strconv.Atoi(input3[i])
		hasil3 += angka3
	}
	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}
