package service

import (
	"Pelatihan9/bangunDatar/shape"
	"fmt"
)

type PerhitunganLuasBangunDatar struct {
	KumpulanBangunDatar []shape.HitungBangunDatar
}

func (p PerhitunganLuasBangunDatar) luasBangunDatar(bd shape.HitungBangunDatar) float64 {
	hasil := bd.Luas()
	//fmt.Println(hasil)
	return hasil
}

func (p PerhitunganLuasBangunDatar) LuasTotal() float64 {
	var hasilTotal float64
	for _, value := range p.KumpulanBangunDatar {
		//hasilTotal += p.luasBangunDatar(value)
		hasilTotal += value.Luas()
	}
	p.cetak(hasilTotal)
	return hasilTotal
}
func (p PerhitunganLuasBangunDatar) cetak(hasil float64) {
	fmt.Println(hasil)
}
