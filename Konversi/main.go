package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	i++
	// Atoi ubah string jadi integer
	// Itoa ubah integer jadi string
	nilaiAngka, _ := strconv.Atoi("42")

	fmt.Println("Hasilnya " + strconv.Itoa(i))
	fmt.Println(nilaiAngka)

	var x int32 = 10
	y := int64(x)
	fmt.Println(y)

	var a float64 = 3.9
	var b int32 = int32(a)
	fmt.Println(b)

	sungguh, _ := strconv.ParseBool("true")
	fmt.Println(sungguh)
}
