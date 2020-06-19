/*gw punya kumpulan Huruf "a", "b", "a", "u", "u", "c", "a"
trus gw mau hitung jumlahnya per kelompok
hasilnya, sbb:
[map[a:3] map[b:1] map[u:2] map[c:1]]

kalau mau coba pake potongan kode dibawah silahkan*/

package main

func main() {
	var huruf = []string{"a", "b", "a", "u", "u", "c", "a"}
	hasil := hitungVarian(huruf)
	fmt.Println(hasil)
}

func hitungVarian(h []string) []map[string]int {
	// koding disini
	hasilHitung := []map[string]int{}
	alphabet := sort.Strings(h)
	jumlah := 1
	for i := 1, i<= len(h); i++ {
		if h[i] == h[i-1] {
			jumlah++
		} else {
			hasilHitung = append(hasilHitung, map[string]int{h[i-1]:jumlah})
			jumlah = 1
		}
	}
	return hasilHitung
}

/*func periksaMap(m []map[string]int, h string) (map[string]int, bool) {
	// koding disini
	for h, jumlah := range m {
		if h ==
	}
}*/
