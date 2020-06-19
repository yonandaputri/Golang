package main

import (
	"Pelatihan9/interfaceHitungTotalLuas/ruangDatar"
	"Pelatihan9/interfaceHitungTotalLuas/service"
)

func main() {
	luasPersegi1 := ruangDatar.Persegi{5}
	luasPersegi2 := ruangDatar.Persegi{10}
	luasSegitiga := ruangDatar.Segitiga{12, 12}
	luasTrapesium := ruangDatar.Trapesium{3, 5, 5}
	sliceLuas := []ruangDatar.HitungBangunDatar{luasPersegi1, luasPersegi2, luasSegitiga, luasTrapesium}

	service.PerhitunganLuasBangunDatar{sliceLuas}.LuasTotal()

}
