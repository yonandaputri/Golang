package utils

// https://tour.golang.org/methods/19
// type error interface {
// 	Error() string
// }

type DibagiNolError struct {
	PesanMasalah string
}

// func (dne DibagiNolError) Error() string {
// 	return dne.PesanMasalah
// }

func (dne DibagiNolError) Error() string {
	if len(dne.PesanMasalah) == 0 {
		return DIVIDED_BY_ZERO
	}
	return dne.PesanMasalah
}
