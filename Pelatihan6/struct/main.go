package main

import "fmt"

// sama kaya class kalau di java. aksesnya pake titik. misal user.name
type user struct {
	name string
	age  int
}

type mobil struct {
	merk  string
	ban   int
	spion int
	warna string
}

func main() {
	// user1 typedatanya user
	/*var user1 user
	user1.name = "budi"
	user1.age = 17

	var user2 = user{"tono", 20}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user1.name)*/

	var allMobil = []mobil{
		{merk: "toyota", spion: 2, ban: 4, warna: "biru"},
		{ban: 2, spion: 2, warna: "merah", merk: "honda"},
		{ban: 6, spion: 4, merk: "suzuki", warna: "kuning"},
		{"bmw", 10, 8, "ijo"},
	}
	for i := 0; i < len(allMobil); i++ {
		if allMobil[i].spion >= 4 {
			fmt.Println(allMobil[i].merk)
		}
	}
	/*for _, elemenMobil := range allMobil {
		if elemenMobil.spion == 4 {
			fmt.Println(elemenMobil.merk)
		}
	}
	fmt.Println(allMobil)*/
}

// alternatif lain
/*var grandMax = mobil{ban: 1, spion: 2, warna: "Merah"}
var toyota = mobil{1, 2, "Merah"}

var banyakMobil = []mobil{grandMax, toyota, {ban: 2, spion: 1, warna: "Biru"}, {10, 2, "Jingga"}}
fmt.Println(banyakMobil)*/
