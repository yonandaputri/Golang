package main

import (
	"Pelatihan9/bangunDatar/service"
	"Pelatihan9/bangunDatar/shape"
)

func main() {
	segitiga := shape.Segitiga{2, 4}
	persegi := shape.Persegi{9}
	persegiPanjang := shape.PersegiPanjang{10, 5}
	trapesium := shape.Trapesium{2, 9, 5}

	rumah := []shape.HitungBangunDatar{segitiga, persegi, persegiPanjang, trapesium}

	(service.PerhitunganLuasBangunDatar{rumah}).LuasTotal()

}
