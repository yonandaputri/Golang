package service

import (
	"Pelatihan9/interfaceHitungTotalLuas/ruangDatar"
	"fmt"
)

type PerhitunganLuasBangunDatar struct {
	KumpulanBangunDatar []ruangDatar.HitungBangunDatar
}

// luasBangunDatar untuk hitung luas per bangun datarnya
func (p PerhitunganLuasBangunDatar) luasBangunDatar(luas ruangDatar.HitungBangunDatar) float64 {
	return luas.HitungLuas()
}

// karena mau dipakai public maka tidak perlu ada parameter
// LuasTotal untuk hitung luas total keseluruhan bangun
func (p PerhitunganLuasBangunDatar) LuasTotal() {
	var totalLuas float64
	for _, bangunDatar := range p.KumpulanBangunDatar {
		totalLuas += p.luasBangunDatar(bangunDatar)
	}
	p.cetak(totalLuas)
}

func (p PerhitunganLuasBangunDatar) cetak(hasilTotal float64) {
	fmt.Println("Hasil luas total : ", hasilTotal)
}
