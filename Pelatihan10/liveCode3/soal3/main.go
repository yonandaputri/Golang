package main

import (
	"fmt"
	"sort"
)

func main() {
	var saldoAwal = 700000
	var idMember string = "123132ssd666"
	// fmt.Println("id:", idMember)
	// fmt.Println("Uang:", saldo)
	var barangSale = map[string]int{
		"Sepatu Ceebook":   1500000,
		"Baju Zoro":        500000,
		"Baju H&N":         250000,
		"Sweater Uniklooh": 175000,
		"Casing Handphone": 50000,
	}

	keys := make([]string, 0, len(barangSale))
	values := make([]int, 0, len(barangSale))

	for k, v := range barangSale {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	jumlahbeli := 0 //isi nilai awal 0
	saldo := 700000
	var barangBeli []string
	for k := 0; k <= len(barangSale)-1; k++ {
		if values[k] <= saldo {
			saldo -= values[k] //ex : 10000 = 10000 - gift yang index of nya dari loop k
			jumlahbeli++       //akan terus bertambah 1 sampai budgetnya abis
			barangBeli = append(barangBeli, keys[k])
		}
	}
	// Output
	var data map[string]interface{}
	data = map[string]interface{}{
		"id":       idMember,
		"uang":     saldoAwal,
		"barang":   barangBeli,
		"sisaUang": saldo,
	}

	fmt.Println(data)
}
