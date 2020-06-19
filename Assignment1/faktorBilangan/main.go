//Viontina Dea IYP

package main

import "fmt"

func main() {
	// input
	var angka int
	fmt.Print("Input angka : ")
	fmt.Scan(&angka)

	for i := angka; i > 0; i-- {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}
