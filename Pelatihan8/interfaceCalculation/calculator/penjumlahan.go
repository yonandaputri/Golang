package calculator

import "Pelatihan8/interfaceCalculation/utils"

type Penjumlahan struct {
	Angka1 int
	Angka2 int
}

func (p Penjumlahan) GetCalculation() (int, error) {
	hasilPenjumlahan := p.Angka1 + p.Angka2
	if p.Angka1 == 0 || p.Angka2 == 0 {
		return 0, utils.PerhitunganError{"Angka Tidak Boleh kosong"}
	} else {
		return hasilPenjumlahan, nil
	}
}
