package utils

// PerhitunganError Message if result negatif
type PerhitunganError struct {
	PesanMasalah string
}

func (err PerhitunganError) Error() string {
	return err.PesanMasalah
}
