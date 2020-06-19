package main

import "fmt"

var penjumlahan = func(angka1, angka2 int) int {
	return angka1 + angka2
}

var pengurangan = func(angka1, angka2 int) int {
	return angka1 - angka2
}

var perkalian = func(angka1, angka2 int) int {
	return angka1 * angka2
}

var pembagian = func(angka1, angka2 int) int {
	return angka1 / angka2
}

func main() {
	f := calculator("/")
	if f != nil {
		fmt.Println(f(4, 2))
	} else {
		fmt.Println("Operator Tidak dikenal")
	}
}

func calculator(operator string) func(int, int) int {
	switch operator {
	case "+":
		return penjumlahan
	case "-":
		return pengurangan
	case "*":
		return perkalian
	case "/":
		return pembagian
	default:
		return nil
	}
}

/*func cetak(angka1, angka2 int) {

}*/

/*func calculator(ops string) func(int, int) int {
	switch ops {
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "-":
		return func(a int, b int) int {
			return a - b
		}

	default:
		return nil
	}
}

func main() {
	f := calculator("+")
	if f != nil {
		fmt.Println(f(1, 3))
	} else {
		fmt.Println("Operator Tidak dikenal")
	}
*/
