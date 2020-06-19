package ruangDatar

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

func (s Segitiga) HitungLuas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}
