package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("")
	fmt.Println(1) //integer
	fmt.Println("")
	fmt.Printf("%T", 1) //cek tipe data
	fmt.Println("")
	fmt.Println(1.3)
	fmt.Println("")
	fmt.Printf("%T", 1.3)
	fmt.Println("")
	fmt.Printf("%v", "HELLO...")

	// boolean (case sensitive)
	fmt.Println("")
	fmt.Println(true)
	fmt.Println(false)

	// modulo (sisa bagi)
	fmt.Println(10 + 1)
	fmt.Println(10 - 1)
	fmt.Println(10 * 2)
	fmt.Println(8 / 3)
	fmt.Println(8.0 / 3.0) //hasilnya float
	fmt.Println(10 % 2)
}
