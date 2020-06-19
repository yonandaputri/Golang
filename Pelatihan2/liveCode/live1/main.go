// Viontina Dea Ivoni YP

package main

import "fmt"

func main() {
	var angka = []int{1, 8}

	if angka[0]%2 == 0 && angka[len(angka)-1]%2 != 0 {
		fmt.Println("One Zombie Down!")
	} else if angka[0]%2 == 0 && angka[len(angka)-1]%2 == 0 {
		fmt.Println("Try Again To Attack")
	} else if angka[0]%2 != 0 && angka[len(angka)-1]%2 != 0 {
		fmt.Println("Try Again To Attack")
	} else {
		fmt.Println("You Are Dead!")
	}
}
