package main

import (
	"fmt"
	"strings"
)

func main() {
	contohKata := "makan"
	var slice []string
	balikKata(&slice, &contohKata)
	sliceToString := strings.Join(slice, "")
	fmt.Println(sliceToString)
}

func balikKata(dibalik *[]string, kata *string) {
	sliceKata := strings.Split(*kata, "")
	for i := len(sliceKata) - 1; i >= 0; i-- {
		if sliceKata[i] != "a" {
			*dibalik = append(*dibalik, sliceKata[i])
		}
	}
}
