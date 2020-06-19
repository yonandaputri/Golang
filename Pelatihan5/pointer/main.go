package main

import "fmt"

func main() {
	var angka int = 20
	//var angka2 int = 50
	fmt.Println("Alamat Memori Angka ", &angka)
	ubahAngka(&angka)
	//ubahAngka(&angka2)
	fmt.Println(angka)
	//fmt.Println(angka2)

	productList := make(map[string]string)
	fmt.Printf("%p\n", &productList)
}

// bertipe data pointer
// tidak lagi membutuhkan return
// alamat memorinya sama dengan yg di main karena ada pointer
// pointer dapat menghemat memori karena tidak perlu menambah variabel lagi untuk return
func ubahAngka(number *int) {
	fmt.Println("Alamat Memory Number ", number)
	fmt.Println(*number)
	// untuk ubah nilai variabelnya ditambah bintang
	*number = 5
}
