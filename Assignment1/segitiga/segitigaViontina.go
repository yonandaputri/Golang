package main

import (
	"fmt"
)

func main() {
	alas := 20
	tinggi := 25

	a := float32(alas)
	t := float32(tinggi)

	// kode disini
	luasSegitiga := (0.5) * a * t
	fmt.Println("luas segitiga dengan alas =", a, "& tinggi =", t, "hasilnya", luasSegitiga)
}
