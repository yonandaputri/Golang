package main

import (
	"Pelatihan8/dictionaryInterface/dictionary"
	"Pelatihan8/dictionaryInterface/utils"
	"fmt"
)

func main() {
	// kalau di dictionarynya pointer disini harus alamat
	// enDict := &dictionary.EnglishDict{}
	// jpDict := &dictionary.KoreaDict{}

	// dictionaries := []dictionary.Dict{enDict, jpDict}

	// for _, dictionary := range dictionaries {
	// 	printGreeting(dictionary)
	// }
	result, err := pembagian(1, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

// tipe data error adalah sebuah interface
func pembagian(num1, num2 int) (int, error) {
	if num2 == 0 {
		// return 0, utils.DibagiNolError{utils.DIVIDED_BY_ZERO}
		return 0, utils.DibagiNolError{}
	}
	result := (num1 / num2)
	return result, nil
}

func printGreeting(d dictionary.Dict) {
	fmt.Println(d.GetMorningGreeting())
}
