package ruangDatar

type Trapesium struct {
	SisiAtas  float64
	SisiBawah float64
	Tinggi    float64
}

func (t Trapesium) HitungLuas() float64 {
	return 0.5 * t.SisiAtas * t.SisiBawah * t.Tinggi
}
