package main

import "fmt"

func main() {
	a := []string{"1", "2"}
	fmt.Printf("Isi A : %v\n", a)
	fmt.Println("len(a)", len(a), "cap(a)", cap(a))
	b := append(a, "3")
	fmt.Printf("Isi B : %v\n", b)
	fmt.Println("len(b)", len(b), "cap(b)", cap(b))

	c := append(b, "4")
	fmt.Printf("Isi C : %v\n", c)
	fmt.Println("len(c)", len(c), "cap(c)", cap(c))
	d := append(b, "5")
	fmt.Printf("Isi D : %v\n", d)
	fmt.Println("len(d)", len(d), "cap(d)", cap(d))
	fmt.Println("Lihatlah isi C menjadi")
	fmt.Printf("Isi C : %v\n", c)
}
