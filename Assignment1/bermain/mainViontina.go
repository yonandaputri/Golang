package main

import "fmt"

func main() {
	var a int = 15
	var b int = 20

	// kode disini
	// a, b = b, a (bisa menggunakan ini)
	/*c := a
	d := b
	a = d
	b = c*/
	c := a
	a = b
	b = c
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
