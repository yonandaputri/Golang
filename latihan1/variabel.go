package main

import (
	"fmt"
)

func main() {
	var namaVariabel string = "budi"
	// var name string = "budi"
	// var tampung string = "budi" (jangan pakai variabel tampung atau variabel yang susah dimengerti orang lain)
	// var name string (nilainya string kosong)
	var age int = 17
	// var age int (nilainya 0)
	// age := 17 (bisa menggunakan seperti ini tipe data mengikuti data yg diinput)
	namaVariabel2 := "budi2"

	age = 20 //redeclared
	//constant
	const firstname string = "tono"
	// firstname = "tini" tidak bisa di assign karena nilai const tidak bisa diubah

	//boolean
	var isEqual = (age == 18)     // penggunaan operator perbandingan
	var umur int                  // umur = 0
	var samaEnggak = (umur == 18) // penggunaan operator perbandingan == != < > <= >=

	//boolean2
	var a = true
	var b = false

	var c = a && b // operator AND
	var d = a || b // operator logika OR
	var e = !b     // operator NOT atau negasi

	fmt.Println(namaVariabel)
	fmt.Println(namaVariabel2)
	fmt.Println(age)
	fmt.Println(firstname)
	fmt.Println(isEqual)
	fmt.Println(samaEnggak)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
