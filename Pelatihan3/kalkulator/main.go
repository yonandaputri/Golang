package main

import "fmt"

func main() {
	hasil := calculator(2, 3, "+")
	fmt.Println(hasil)
}

func calculator(bil1, bil2 int, operator string) int {
	var answer int
	if operator == "+" {
		answer = bil1 + bil2
	} else if operator == "-" {
		answer = bil1 - bil2
	} else if operator == "*" {
		answer = bil1 * bil2
	} else if operator == "/" {
		answer = bil1 / bil2
	}
	return answer
}
