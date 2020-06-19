package main

import "fmt"

func main() {
	var alamatIntLangsung = new(int)
	fmt.Println(alamatIntLangsung)
	*alamatIntLangsung = 2
	// untuk print alamat pointer
	fmt.Println(&alamatIntLangsung)
	// untuk print value
	fmt.Println(*alamatIntLangsung)
	//ubahAngka(alamatIntLangsung)
}

/*func ubahAngka(number *int) {
	*number = 2
	fmt.Println(*number)
}*/
