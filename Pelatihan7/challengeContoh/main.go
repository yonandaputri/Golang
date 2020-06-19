package main

import "fmt"

type mahasiswa struct {
	nama    string
	jurusan string
}

func (m *mahasiswa) toString() string {
	sNama := fmt.Sprintf("%-15s %s\n", "Nama Lengkap", m.nama)
	sJurusan := fmt.Sprintf("%-15s %s\n", "Jurusan", m.jurusan)
	return sNama + sJurusan
}

func newMahasiswa(nama, jurusan string) *mahasiswa {
	return &mahasiswa{
		nama:    nama,
		jurusan: jurusan,
	}
}

func main() {
	mahasiswa1 := newMahasiswa("Sasa", "Matematika")
	mahasiswa2 := newMahasiswa("Apin", "Kedokteran Binatang")
	fmt.Println(mahasiswa1.toString())
	fmt.Println(mahasiswa2.toString())
}
