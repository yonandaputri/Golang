package main

import "fmt"

func main() {
	tulisan := "Hai"
	i := 0
	for i < 5 {
		fmt.Println(tulisan)
	}

	/*angka := 1
	if angka == 1 {
		fmt.Println("angka 1")
	} else if angka == 2 {
		fmt.Println("bukan angka 1")
	} else {
		fmt.Println("bukan apa2")
	}*/

	//switch
	angka := 0
	switch angka {
	case 0:
		fmt.Println("nol")
	default:
		fmt.Println("Salah")
	}
}
