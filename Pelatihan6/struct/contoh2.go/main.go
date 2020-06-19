//
package main

import "fmt"

type siswa struct {
	name  string
	grade string
}

type kilogram int

type stringku [][]string // slice custom

type crazyVariabel []map[string]map[string]string

func main() {
	var user siswa
	fmt.Println(user.name)
	var garam kilogram = 1
	fmt.Println(garam)
}
